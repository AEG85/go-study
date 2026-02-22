package pressure

type Coordinat struct {
	X float64
	Y float64
}

var coordinates = map[int]Coordinat{
	1: {X: 12, Y: 15},
	2: {X: 13, Y: 16},
	3: {X: 33, Y: 12.7},
	4: {X: 2, Y: 17},
	5: {X: 66, Y: 77},
}

func SensorCoordinates(sensorIndex int) Coordinat {
	coords, ok := coordinates[sensorIndex]
	if !ok {
		return Coordinat{X: 1, Y: 2}
	}
	return coords
}
