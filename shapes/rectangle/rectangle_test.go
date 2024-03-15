package shapes

import "testing"

type RectangleTestCase struct {
	TestCaseName string
	Length       float64
	Breadth      float64
	Want         float64
	ErrWant      error
}

func TestRectangleAreaWithLengthBreadthIsPositive(t *testing.T) {
	want := 6.0
	rectangle, err := NewRectangle(2.0, 3.0)
	got := rectangle.Area()
	if want != got {
		t.Errorf("Error: want %f got %f", want, got)
	}
	if err != nil {
		t.Errorf("Error: error want %s error got %s", "nil", err.Error())
	}
}

func TestRectangleAreaWithLengthIsNegative(t *testing.T) {
	want := -1.0
	rectangle, err := NewRectangle(-2.0, 3.0)
	got := rectangle.Area()
	if want != got {
		t.Errorf("Error: want %f got %f", want, got)
	}
	if err != ErrLengthIsNegative {
		t.Errorf("Error: error want %s error got %s", ErrLengthIsNegative.Error(), err.Error())
	}
}
func TestRectangleAreaWithBreadthIsNegative(t *testing.T) {
	want := -1.0
	rectangle, err := NewRectangle(2.0, -3.0)
	got := rectangle.Area()
	if want != got {
		t.Errorf("Error: want %f got %f", want, got)
	}
	if err != ErrBreadthIsNegative {
		t.Errorf("Error: error want %s error got %s", ErrBreadthIsNegative.Error(), err.Error())
	}
}

func TestRectangleAreaWithLengthAndBreadthIsNegative(t *testing.T) {
	want := -1.0
	rectangle, err := NewRectangle(-2.0, -3.0)
	got := rectangle.Area()
	if want != got {
		t.Errorf("Error: want %f got %f", want, got)
	}
	if err != ErrLengthAndBreadthIsNegative {
		t.Errorf("Error: error want %s error got %s", ErrLengthAndBreadthIsNegative.Error(), err.Error())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	testCases := []RectangleTestCase{
		{
			TestCaseName: "TestRectanglePerimeterWithLengthBreadthIsPositive",
			Length:       3.0,
			Breadth:      2.0,
			ErrWant:      nil,
			Want:         10.0,
		},
		{
			TestCaseName: "TestRectanglePerimeterWithLengthBreadthIsNegative",
			Length:       -3.0,
			Breadth:      -2.0,
			ErrWant:      ErrLengthAndBreadthIsNegative,
			Want:         0.0,
		},
		{
			TestCaseName: "TestRectanglePerimeterWithLengthIsNegative",
			Length:       3.0,
			Breadth:      2.0,
			ErrWant:      ErrLengthIsNegative,
			Want:         0.0,
		},
		{
			TestCaseName: "TestRectanglePerimeterWithBreadthIsNegative",
			Length:       3.0,
			Breadth:      2.0,
			ErrWant:      ErrBreadthIsNegative,
			Want:         0.0,
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.TestCaseName,
			func(t *testing.T) {

				rectangle, errGot := NewRectangle(testCase.Length, testCase.Breadth)
				if errGot != testCase.ErrWant {
					t.Errorf("Error: error want %s, error got %s", errGot, testCase.ErrWant)
					return
				}
				got := rectangle.Area()

				if got != testCase.Want {
					t.Errorf("Error: want %f, got %f", testCase.Want, got)
				}

			},
		)
	}
}
