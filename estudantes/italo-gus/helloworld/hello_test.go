package main // TDD com subtestes e IF

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		/* Run(nome string, f func(t *T)) bool
		Run executa f como um subteste de t, chamado nome.
		Ele executa f em uma goroutine separada e bloqueia até que f retorne ou chame t.Parallel
		para se tornar um teste paralelo. Run relata se f foi bem-sucedido
		(ou pelo menos não falhou antes de chamar t.Parallel).
		Run pode ser chamado simultaneamente de várias goroutines,
		mas todas essas chamadas devem retornar antes que a função de teste externa para t retorne.
		https://pkg.go.dev/testing#T.Run
		https://go.dev/doc/effective_go#goroutines
		*/

		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

/*
Para funções auxiliares, é uma boa ideia aceitar uma testing.TB
que é uma interface que *testing.Te *testing.B
ambas satisfazem, para que você possa chamar funções auxiliares de um teste ou de um benchmark
*/
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	/*	Helper marca a função de chamada como uma função de teste auxiliar.
		Ao imprimir informações de arquivo e linha, essa função será ignorada.
		Helper pode ser chamado simultaneamente de várias goroutines.
		https://pkg.go.dev/testing#T.Helper
	*/

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
