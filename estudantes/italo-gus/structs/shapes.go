package structs // struct - estrutura

/* Declarando um tipo baseado na criação de uma estrutura

type Nome_do_tipo struct {
campo1_da_estrutura typo_de_dados_do_campo1_da_estrutura
campon_da_estrutura typo_de_dados_do_campon_da_estrutura   }

https://go.dev/ref/spec#Types
https://go.dev/ref/spec#Struct_types
*/
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

// acessar campos de uma estrutura variáve.campo_da_estrutura
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
