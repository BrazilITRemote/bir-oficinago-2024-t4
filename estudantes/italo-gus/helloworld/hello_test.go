package main // TDD com subtestes go test -v para expor relatório dos testes

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("em espanhol", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("em francês", func(t *testing.T) {
		got := Hello("Francis", "French")
		want := "Bonjour, Francis"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	/*
		Helper marca a função de chamada como uma função de teste auxiliar.
		Ao imprimir informações de arquivo e linha, essa função será ignorada.
		Helper pode ser chamado simultaneamente de várias goroutines.
		https://pkg.go.dev/testing#T.Helper
	*/
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
no terminal go test -v para expor relatório dos testes
=== RUN   TestHello
=== RUN   TestHello/diz_olá_para_as_pessoas
=== RUN   TestHello/'Mundo'_como_padrão_para_'string'_vazia
=== RUN   TestHello/em_espanhol
=== RUN   TestHello/em_francês
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/diz_olá_para_as_pessoas (0.00s)
    --- PASS: TestHello/'Mundo'_como_padrão_para_'string'_vazia (0.00s)
    --- PASS: TestHello/em_espanhol (0.00s)
    --- PASS: TestHello/em_francês (0.00s)
PASS
ok      learn-go-with-tests/helloworld_07
*/
