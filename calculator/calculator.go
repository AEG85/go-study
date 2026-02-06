package calculator

import (
	"errors"
	"strconv"
)

type Calculator struct {
	X float64
	Y float64
}

func NewCalculator(x float64, y float64) Calculator {
	return Calculator{
		X: x,
		Y: y,
	}
}

func (c *Calculator) Summ() (float64, error) {
	if (c.X > 1000 || c.X < -1000) || (c.Y > 1000 || c.Y < -1000) {
		error := errors.New("Аргументы не должны быть > 1000 и < -1000. " + MakeStringArgs(c.X, c.Y))
		return 0, error
	}
	return c.X + c.Y, nil
}

func (c *Calculator) Subtract() (float64, error) {
	if (c.X > 1000 || c.X < -1000) || (c.Y > 1000 || c.Y < -1000) {
		error := errors.New("Аргументы не должны быть > 1000 и < -1000. " + MakeStringArgs(c.X, c.Y))
		return 0, error
	}
	return c.X - c.Y, nil
}

func (c *Calculator) Multyply() (float64, error) {
	if (c.X > 1000 || c.X < -1000) || (c.Y > 1000 || c.Y < -1000) {
		error := errors.New("Аргументы не должны быть > 1000 и < -1000. " + MakeStringArgs(c.X, c.Y))
		return 0, error
	}
	return c.X * c.Y, nil
}

func (c *Calculator) Devide() (float64, error) {
	if (c.X > 1000 || c.X < -1000) || (c.Y > 1000 || c.Y < -1000) {
		error := errors.New("Аргументы не должны быть > 1000 и < -1000. " + MakeStringArgs(c.X, c.Y))
		return 0, error
	}
	if c.Y == 0 {
		error := errors.New("Деление на ноль запрещено. " + MakeStringArgs(c.X, c.Y))
		return 0, error
	}
	return c.X / c.Y, nil
}

func MakeStringArgs(x float64, y float64) string {
	return "X: " + strconv.FormatFloat(x, 'f', -1, 64) + ", Y: " + strconv.FormatFloat(y, 'f', -1, 64)
}
