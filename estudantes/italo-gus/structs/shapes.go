package structs // TableDrivenTests

import (
	"math"
)

/* math O pacote math fornece constantes básicas e funções matemáticas.
Este pacote não garante resultados idênticos em bits em todas as arquiteturas.
https://pkg.go.dev/math@go1.23.4
*/

/*
Ao incorporar interfaces, métodos com os mesmos nomes devem ter assinaturas idênticas .
https://go.dev/ref/spec#Interface_types
https://go.dev/doc/effective_go#interface_methods
*/
// Criando tipo Shape, que é uma interface { definindo o método Area com tipo de retorno float64}
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

/*
func Perimeter(width, height float64) float64 {
	//return 0
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}
*/

/*
	É uma convenção em Go que a variável receptora seja a primeira letra do tipo.
	https://go.dev/ref/spec#Identifiers
	https://google.github.io/styleguide/go/decisions#receiver-names
	https://google.github.io/styleguide/go/decisions.html#single-letter-variable-names
*/

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

/*
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
*/

func (r Rectangle) Area() float64 { //----   Método definido pela interface
	return r.Width * r.Height
}

/*
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
*/

func (c Circle) Perimeter() float64 {
	return math.Pi * (2 * c.Radius)
}

func (c Circle) Area() float64 { //---           Método definido pela interface
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64 { //---          Método definido pela interface
	return (t.Base * t.Height) * 0.5
}
