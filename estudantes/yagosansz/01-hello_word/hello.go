package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	portuguese            = "Portuguese"
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Oi, "
)

// In Go, public functions start with a capital letter, whereas
// private functions start with lowercase letter.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// (prefix string) is a named return.
// It creates the prefix variable and assigns it a "zero" value.
// For strings the "zero" value is an empty string.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
