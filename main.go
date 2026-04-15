package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Strings = []string{}
var mu sync.RWMutex = sync.RWMutex{}

const (
	CantWriteResponse = "Не получилось записать ответ"
	InvalidJson       = "Не получилось преобразовать данные в json формат"
)

func main() {
	r := chi.NewRouter()
	logger, loggerClose, err := NewLogger("DEBUG")
	if err != nil {
		panic(err)
	}
	defer loggerClose()

	r.Route("/strings", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			getStrings(w, r, logger)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			addString(w, r, logger)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getStrings(w, r, logger)
	})

	http.ListenAndServe(":9091", r)
}

func addString(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Произошла ошибка при чтении: ", zap.Error(err))
		return
	}
	startTime := time.Now()
	logger.Info("Входящий запрос:",
		zap.String("method", r.Method),
		zap.String("endpoint", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("request_body", string(body)),
		zap.Time("request_time", startTime),
	)

	mu.Lock()
	defer mu.Unlock()

	Strings = append(Strings, string(body))

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Элемент добавлен"))
	if err != nil {
		logger.Error("Не получилось отправить ответ", zap.Error(err))
	}

	duration := time.Since(startTime)
	logger.Info("Добавлена новая строка",
		zap.String("added_value", string(body)),
		zap.Int("slice_length", len(Strings)),
		zap.Duration("processing_time", duration),
	)
}

func getStrings(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Произошла ошибка при чтении: ", zap.Error(err))
		return
	}
	startTime := time.Now()

	logger.Info("Входящий запрос:",
		zap.String("method", r.Method),
		zap.String("endpoint", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("request_body", string(body)),
		zap.Time("request_time", startTime),
	)
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Strings); err != nil {
		logger.Error("Ошибка кодирования JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	duration := time.Since(startTime)
	logger.Info("Строки переданы успешно!",
		zap.Duration("processing_time", duration),
	)
}

func NewLogger(loggerLevel string) (*zap.Logger, func() error, error) {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(loggerLevel)); err != nil {
		return nil, nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join("logs", fmt.Sprintf("%s.log", timestamp))

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	encoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl),
		zapcore.NewCore(encoder, zapcore.AddSync(logFile), lvl),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return logger, logFile.Close, nil
}
