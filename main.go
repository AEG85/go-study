package main

import (
	"study/geometry"
	"study/geometry/figures"

	"github.com/k0kubun/pp"
)

func main() {

	params := figures.FigureParams{
		Width:  1,
		Height: 2,
		Radius: 1,
	}

	figure := figures.NewSquare(params)

	geometryModule := geometry.GeometryModule{
		Figure: figure,
	}
	pp.Println("Площадь фигуры: ", geometryModule.Figure.GetArea())
	pp.Println("Периметр фигуры: ", geometryModule.Figure.GetPerimetr())
	pp.Println("Длина окружности: ", geometryModule.Figure.GetCircleLength())
}
