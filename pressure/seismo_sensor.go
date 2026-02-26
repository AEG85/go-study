package pressure

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func SeismoSensor(ctx context.Context, meteoTransfer chan<- string) {
	sensorNumber := 0
	for {
		sensorNumber++
		select {
		case <-ctx.Done():
			fmt.Println("Отменяем сбор данных датчика Сейсмоактивности!")
			return
		case <-time.After(1 * time.Second):
			coordinates := SensorCoordinates(sensorNumber)
			sensorInfo := "Я датчик сейсмоактивности #" + strconv.Itoa(sensorNumber) + ". S " + strconv.Itoa(sensorNumber*10) + ". Мои координаты: X:" + strconv.FormatFloat(coordinates.X, 'f', -1, 64) + " Y:" + strconv.FormatFloat(coordinates.Y, 'f', -1, 64)
			meteoTransfer <- sensorInfo
		}
	}
}
