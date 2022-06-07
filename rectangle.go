package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	//it should be width instead of weight but Autocode tests can't be modified :(
	Height, Weight float64
}

func (f Rectangle) CalcPerimeter() float64 {
	return (f.Height + f.Weight) * 2
}

func (f Rectangle) CalcArea() float64 {
	return f.Height * f.Weight
}
