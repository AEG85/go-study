package geometry

type GeometryFigure interface {
	GetArea() float64
	GetPerimetr() float64
	GetCircleLength() float64
}

type GeometryModule struct {
	Figure GeometryFigure
}
