package pressure

import (
	"fmt"
	"strconv"
)

type Coordinat struct {
	x float64
	y float64
}

var coordinates = map[int]Coordinat{
	1: {X: 12, Y: 15},
	2: {X: 13, Y: 16},
	3: {X: 33, Y: 12.7},
	4: {X: 2, Y: 17},
	5: {X: 66, Y: 77},
}

func PressureSensor(pressureTransfer chan int64, sensorNumber int) {
	for {
		fmt.Println("Я датчик #", strconv.Itoa(sensorNumber), ". Давление ", strconv.Itoa)
	}
}

func SensorCoordinates(sensorIndex int) Coordinat {
	coords, ok := coordinates[sensorIndex]
	if !ok {
		return Coordinat{X: 1, Y: 2}
	}
	return coords
}
