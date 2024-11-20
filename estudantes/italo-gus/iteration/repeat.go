package iteration // Benchmarks em Go
import "fmt"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string = ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatCount(character string, repeatCountVar int) string {
	var repeated string = ""
	for i := 0; i < repeatCountVar; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat() bool {
	fmt.Println("Repeat('a')")
	fmt.Println(Repeat("a"))
	fmt.Println("RepeatCount('a',3)")
	fmt.Println(RepeatCount("a", 3))
	return true
}

/*
Por que o Go não suporta sobrecarga de métodos e operadores?
https://go.dev/doc/faq#overloading
*/
