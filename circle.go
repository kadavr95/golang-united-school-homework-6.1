package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (f Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * f.Radius
}

func (f Circle) CalcArea() float64 {
	return math.Pi * math.Pow(f.Radius, 2)
}
