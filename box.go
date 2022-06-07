package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorCapacityOverflow = errors.New("the box is already full")
	errorShapeNotFound    = errors.New("there is no shape at this index")
	errorNoSuchShapes     = errors.New("there are no circles in the box")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         make([]Shape, 0),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes)+1 <= b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("an error occurred: %w", errorCapacityOverflow)
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) > i {
		return b.shapes[i], nil
	}
	return nil, fmt.Errorf("an error occurred: %w", errorShapeNotFound)
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) > i {
		extractedShape := b.shapes[i]
		// could be faster without order preservation
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return extractedShape, nil
	}
	return nil, fmt.Errorf("an error occurred: %w", errorShapeNotFound)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) > i {
		extractedShape := b.shapes[i]
		b.shapes[i] = shape
		return extractedShape, nil
	}
	return nil, fmt.Errorf("an error occurred: %w", errorShapeNotFound)
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, f := range b.shapes {
		sum += f.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := 0.0
	for _, f := range b.shapes {
		sum += f.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	flag := false
	for i := len(b.shapes) - 1; i >= 0; i-- {
		switch b.shapes[i].(type) {
		case *Circle:
			_, _ = b.ExtractByIndex(i)
			flag = true
		}
	}
	if flag {
		return nil
	}
	return fmt.Errorf("an error occurred: %w", errorNoSuchShapes)
}
