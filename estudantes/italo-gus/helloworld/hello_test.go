package main // TDD

/* testing O teste de pacote fornece suporte para teste automatizado de pacotes Go. Ele é destinado a ser usado em conjunto com o comando "go test", que automatiza a execução de qualquer função do formulário.
https://pkg.go.dev/testing@go1.23.3

O nome do arquivo deve ser formado pelo arquivo a ser testado adicionado _test
*/

import (
	"testing"
)

/*
Função para realizar o teste de ser formado por Test adicionado o nome da função a ser testada
necessitando do parâmetro t *testing.T
*/
func TestHello(t *testing.T) {
	got := Hello()         // variáve que recebe o retorno do resultado da função a ser testada
	want := "Hello, World" // variáve que armazena o valor esperado de retorno para comparação do teste

	if got != want { // verifica se há diferença entre as variáveis o resultado da função e o esperado
		t.Errorf("got %q want %q", got, want)
		/* Método Errorf de  t , que imprimirá uma mensagem no caso de falha no teste.
		O f significa format, o que nos permite construir uma string com valores inseridos nos valores do placeholder %q. */
	}

	/* no terminal go test
		** alterado para gerar falha
		--- FAIL: TestHello (0.00s)
	    	hello_test.go:22: got "Hello, world" want "Hello, World"
		FAIL
		exit status 1
		FAIL    learn-go-with-tests/helloworld_03 */

}
