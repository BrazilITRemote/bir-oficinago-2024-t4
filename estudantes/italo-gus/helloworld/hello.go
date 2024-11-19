package main // switch, declaração de variáveis (constantes) em bloco

import "fmt"

/*	declaração de variáveis (constantes) em bloco
	https://go.dev/ref/spec#Variable_declarations */

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	/*	switch alvo { case 0: realize case 1: realize default: realize }
		https://go.dev/doc/effective_go#if
		https://go.dev/doc/effective_go#switch
		https://go.dev/ref/spec#If_statements
		https://go.dev/ref/spec#Switch_statements */

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
