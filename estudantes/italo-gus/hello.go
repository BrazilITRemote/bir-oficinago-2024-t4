package main // Definição do pacote

/* fmt O pacote fmt implementa E/S formatadas com funções análogas a printf e scanf do C. Os 'verbos' de formato são derivados do C, mas são mais simples.
https://pkg.go.dev/fmt@go1.23.3
*/
import "fmt" // Importar Biblioteca

// Criação de Função func nome_da_Função () tipo de retorno {} sem parâmetros
func Hello() string {
	return "Hello, world" // retorno da função
}

// Criação de Função func nome_da_Função () {} sem parâmetros
func main() {
	fmt.Println(Hello()) // imprimir em tela o retorno da função
}
