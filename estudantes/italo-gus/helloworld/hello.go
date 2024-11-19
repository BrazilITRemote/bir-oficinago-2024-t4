package main

import "fmt"

// Aplicando S.O.L.I.D: S — Single Responsiblity Principle (Princípio da responsabilidade única)
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
