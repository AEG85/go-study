package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

	storage := make(map[string]string)
	mtx := sync.RWMutex{}
	wg := sync.WaitGroup{}

	start := time.Now()

	for i := range 100_000 {
		wg.Add(1)
		keyVal := strconv.Itoa(i)
		go AddString(&storage, &wg, &mtx, keyVal, keyVal)
	}

	for i := range 100_000 {
		wg.Add(1)
		key := strconv.Itoa(i)
		go GetString(&storage, &wg, &mtx, key)
	}

	wg.Wait()

	fmt.Println("Количество эл-ов в мапе: ", len(storage))
	fmt.Println("Время работы скрипта: ", time.Since(start))

}

func AddString(storage *map[string]string, wg *sync.WaitGroup, mtx *sync.RWMutex, strKey string, strVal string) bool {
	defer wg.Done()

	mtx.RLock()
	_, ok := (*storage)[strKey]
	mtx.RUnlock()
	if ok {
		return false
	}
	mtx.Lock()
	(*storage)[strKey] = strVal
	mtx.Unlock()
	return true
}

func GetString(storage *map[string]string, wg *sync.WaitGroup, mtx *sync.RWMutex, strKey string) (string, error) {
	defer wg.Done()

	mtx.RLock()
	str, ok := (*storage)[strKey]
	mtx.RUnlock()

	if !ok {
		error := errors.New("Нет такой строки")
		return "", error
	} else {
		return str, nil
	}
}
