package shapes

import "errors"

var ErrLengthIsNegative = errors.New("length is negative")
var ErrLengthIsZero = errors.New("length is zero")

type Square struct {
	Length float64
}

func NewSquare(length float64) (*Square, error) {
	if length < 0 {
		return nil, ErrLengthIsNegative
	} else if length == 0 {
		return nil, ErrLengthIsZero
	}
	return &Square{
		Length: length,
	}, nil
}
func (square *Square) Area() float64 {
	return square.Length * square.Length
}

func (square *Square) Perimeter() float64 {
	return 4.0 * square.Length
}
