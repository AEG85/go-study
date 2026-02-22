package pressure

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func PressureSensor(ctx context.Context, wg *sync.WaitGroup, meteoTransfer chan<- string) {
	defer wg.Done()
	sensorNumber := 0
	for {
		sensorNumber++
		select {
		case <-ctx.Done():
			fmt.Println("Отменяем сбор данных датчика Давления!")
			return
		case <-time.After(1 * time.Second):
			coordinates := SensorCoordinates(sensorNumber)
			sensorInfo := "Я датчик давления #" + strconv.Itoa(sensorNumber) + ". P " + strconv.Itoa(sensorNumber*10) + ". Мои координаты: X:" + strconv.FormatFloat(coordinates.X, 'f', -1, 64) + " Y:" + strconv.FormatFloat(coordinates.Y, 'f', -1, 64)
			meteoTransfer <- sensorInfo
		}
	}
}
