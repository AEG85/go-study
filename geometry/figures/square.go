package figures

type Square struct {
	Params FigureParams
}

func NewSquare(params FigureParams) Square {
	if params.Width > 0 {
		return Square{
			Params: params,
		}
	} else {
		return Square{
			Params: FigureParams{},
		}
	}
}

func (s Square) GetArea() float64 {
	if s.Params.Width > 0 {
		area := s.Params.Width * s.Params.Width
		return area
	}
	return 0
}

func (s Square) GetPerimetr() float64 {
	if s.Params.Width > 0 {
		perimetr := (s.Params.Width) * 4
		return perimetr
	}
	return 0
}

func (s Square) GetCircleLength() float64 {
	return 0
}
