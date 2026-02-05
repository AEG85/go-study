package figures

import "math"

type Circle struct {
	Params FigureParams
}

func NewCircle(params FigureParams) Circle {
	if params.Radius > 0 {
		return Circle{
			Params: params,
		}
	} else {
		return Circle{
			Params: FigureParams{},
		}
	}
}

func (c Circle) GetArea() float64 {
	if c.Params.Radius > 0 {
		area := math.Pi * c.Params.Radius * c.Params.Radius
		return area
	}
	return 0
}

func (c Circle) GetPerimetr() float64 {
	return 0
}

func (c Circle) GetCircleLength() float64 {
	if c.Params.Radius > 0 {
		length := 2 * math.Pi * c.Params.Radius
		return length
	}
	return 0
}
