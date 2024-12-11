package structs_methods_interfaces

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	fmt.Printf("BIR - type of r: %s\n", reflect.TypeOf(r))
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	fmt.Printf("BIR - type of c: %s\n", reflect.TypeOf(c))
	return math.Pi * c.Radius * c.Radius
}
