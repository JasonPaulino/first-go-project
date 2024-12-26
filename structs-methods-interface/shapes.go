package main

// Structs & Interfaces
type Rectangle struct {
	Width  float64
	Length float64
}

type Triangle struct {
	Base   float64
	Height float64
	Region float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Reactangle Behaviors
func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Width
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Length + rectangle.Width)
}


// Triangle Behaviors
func (triangle Triangle) Area() float64 {
	return 0.5 * (triangle.Base * triangle.Height)
}

func (triangle Triangle) Perimeter() float64 {
	return triangle.Base + triangle.Region + triangle.Height
}
