package figures

type Rectangular struct {
	Params FigureParams
}

func NewRectangular(params FigureParams) Rectangular {
	if params.Width > 0 && params.Height > 0 {
		return Rectangular{
			Params: params,
		}
	} else {
		return Rectangular{
			Params: FigureParams{},
		}
	}
}

func (r Rectangular) GetArea() float64 {
	if r.Params.Width > 0 {
		area := r.Params.Width * r.Params.Width
		return area
	}
	return 0
}

func (r Rectangular) GetPerimetr() float64 {
	if r.Params.Width > 0 {
		perimetr := (r.Params.Width + r.Params.Height) * 2
		return perimetr
	}
	return 0
}

func (r Rectangular) GetCircleLength() float64 {
	return 0
}
