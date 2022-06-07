package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (f Triangle) CalcPerimeter() float64 {
	return f.Side * 3
}

func (f Triangle) CalcArea() float64 {
	return math.Pow(f.Side, 2) * math.Sqrt(3) / 4
}
