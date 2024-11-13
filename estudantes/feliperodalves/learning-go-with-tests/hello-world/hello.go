package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHellpPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func gettingPrefix(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = spanishHellpPrefix
		case french:
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Felipe", "Spanish"))
	fmt.Println(Hello("Felipe", "French"))
	fmt.Println(Hello("Felipe", "Portuguese"))
}