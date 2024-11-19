package main // tdd com argumento

import "fmt"

// Criação da Função Hello (com parâmetros ) tipo de retorno string {}
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world")) // imprimir em tela (stdout [saída padrão do terminal]) o retorno da função Hello passando argumento
}
