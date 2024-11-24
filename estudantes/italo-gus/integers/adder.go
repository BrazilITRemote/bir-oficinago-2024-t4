package integers // Go Doc
import "fmt"

// Adiciona recebe dois inteiros e retorna um inteiro que é a soma deles
func Add(x, y int) int {
	return x + y
}

// Exemplo de teste da função ADD(1,5) e imprime o retorno no terminal
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

}

/*
no terminal godoc -http=:6060
no navegador http://localhost:6060/pkg/learn-go-with-tests/integers_02/
o terminal fica bloqueado, no menu terminal selecione new terminal
*/
