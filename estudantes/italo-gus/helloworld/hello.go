package main // Aplicando S.O.L.I.D, E parâmetros de resultado nomeados

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Aplicando S.O.L.I.D: S — Single Responsiblity Principle (Princípio da responsabilidade única)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

/* Parâmetros de resultado nomeados
https://go.dev/wiki/CodeReviewComments#named-result-parameters
https://go.dev/doc/effective_go#named-results*/

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
