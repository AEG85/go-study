package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	wg := &sync.WaitGroup{}

	countWorkers, err := strconv.Atoi(os.Getenv("WORKERS_COUNT"))
	if err != nil {
		fmt.Println("Не корректное число рабочих:", err)
		return
	}

	for workerInderx := range countWorkers {
		wg.Add(1)
		go work(ctx, wg, workerInderx+1)
	}

	wg.Wait()
	fmt.Println("Все рабочие завершили работу")
}

func work(ctx context.Context, wg *sync.WaitGroup, workerNumber int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Работник №", workerNumber, "закончил свою работу!")
			return
		default:
			<-time.After(2 * time.Second)
			fmt.Println("Работник №", workerNumber, "работает...")
		}
	}
}
