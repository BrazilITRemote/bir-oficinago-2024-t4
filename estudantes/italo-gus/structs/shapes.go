package structs // Declarações de métodos

type Rectangle struct {
	Width  float64
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

/* Declarações de métodos
func (nome_do_receptor tipo_do_receptor) Nome_do_método(parâmetros) tipo_do_retorno { bloco_do_método }
https://go.dev/ref/spec#Method_declarations
https://go.dev/doc/effective_go#interface_methods
*/

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

/*
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
*/

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

/*
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
*/
