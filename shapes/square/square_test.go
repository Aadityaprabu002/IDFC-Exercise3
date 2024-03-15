package shapes

import (
	"testing"
)

type SquareTestCase struct {
	TestCaseName string
	Length       float64
	Want         float64
	ErrWant      error
}

func TestSquareArea(t *testing.T) {

	testCases := []SquareTestCase{
		{
			TestCaseName: "TestSquareAreaWithLengthIsPositive",
			Length:       3.0,
			ErrWant:      nil,
			Want:         9.0,
		},
		{
			TestCaseName: "TestSquareAreaWithLengthIsNegative",
			Length:       -3.0,
			ErrWant:      ErrLengthIsNegative,
			Want:         0.0,
		},
		{
			TestCaseName: "TestSquareAreaWithLengthIsNegative",
			Length:       0.0,
			ErrWant:      ErrLengthIsZero,
			Want:         0.0,
		},
	}
	for _, testCase := range testCases {
		t.Run(
			testCase.TestCaseName,
			func(t *testing.T) {

				square, errGot := NewSquare(testCase.Length)

				if errGot != testCase.ErrWant {
					t.Errorf("Error: error want %s, error got %s", errGot, testCase.ErrWant)
					return
				}

				if square != nil {
					got := square.Area()
					if got != testCase.Want {
						t.Errorf("Error: want %f, got %f", testCase.Want, got)
					}
				}
			},
		)
	}

}

func TestSquarePerimeter(t *testing.T) {
	testCases := []SquareTestCase{
		{
			TestCaseName: "TestSquarePerimeterWithLengthIsPositive",
			Length:       3.0,
			ErrWant:      nil,
			Want:         12.0,
		},
		{
			TestCaseName: "TestSquarePerimeterWithLengthIsNegative",
			Length:       -3.0,
			ErrWant:      ErrLengthIsNegative,
			Want:         0.0,
		},
		{
			TestCaseName: "TestSquarePerimeterWithLengthIsNegative",
			Length:       0.0,
			ErrWant:      ErrLengthIsZero,
			Want:         0.0,
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.TestCaseName,
			func(t *testing.T) {

				square, errGot := NewSquare(testCase.Length)

				if errGot != testCase.ErrWant {
					t.Errorf("Error: error want %s, error got %s", errGot, testCase.ErrWant)
					return
				}
				if square != nil {
					got := square.Perimeter()
					if got != testCase.Want {
						t.Errorf("Error: want %f, got %f", testCase.Want, got)
					}
				}

			},
		)
	}
}
