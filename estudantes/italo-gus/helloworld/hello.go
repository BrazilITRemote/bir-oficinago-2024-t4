package main // TDD com subtestes e IF

import "fmt"

const englishHelloPrefix = "Hello, " // Definindo constante

func Hello(name string) string {
	if name == "" { // teste de condição ( se condição { realize } )
		name = "World"
	}
	/* https://go.dev/doc/effective_go#if
	https://go.dev/ref/spec#If_statements */

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
