package main // const

import "fmt"

const englishHelloPrefix = "Hello, " // Definindo constante

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
