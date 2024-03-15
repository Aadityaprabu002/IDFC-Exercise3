package shapes

import "errors"

var ErrLengthIsNegative = errors.New("length is negative")
var ErrBreadthIsNegative = errors.New("breadth is negative")
var ErrLengthAndBreadthIsNegative = errors.New("length and breadth is negative")

type Rectangle struct {
	Length  float64
	Breadth float64
}

func NewRectangle(length, breadth float64) (*Rectangle, error) {
	if length < 0 && breadth < 0 {
		return nil, ErrLengthAndBreadthIsNegative
	} else if breadth < 0 {
		return nil, ErrBreadthIsNegative
	} else if length < 0 {
		return nil, ErrLengthIsNegative
	}
	return &Rectangle{
		Length:  length,
		Breadth: breadth,
	}, nil
}
func (rectangle *Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Breadth
}

func (rectangle *Rectangle) Perimeter() float64 {
	return 2.0 * (rectangle.Length + rectangle.Breadth)
}
