package main

import "fmt"

// func Hello() string {
// 	return "Hello, world"
// }

// func main() {
// 	fmt.Println(Hello())
// }

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
