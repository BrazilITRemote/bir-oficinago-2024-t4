package structs_methods_interfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}
