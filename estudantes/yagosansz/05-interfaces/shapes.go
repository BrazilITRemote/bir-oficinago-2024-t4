package main

import "math"

// Using a struct helps us to convey our intentions
// more clearly!
type Rectangle struct {
	Width  float64
	Height float64
}

// func (receiverName ReceiverType) MethodName(args)
//
// It's a convention that the receiverName is the first letter
// of the ReceiverType
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
