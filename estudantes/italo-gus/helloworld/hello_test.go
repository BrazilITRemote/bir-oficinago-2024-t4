package main // Aplicando S.O.L.I.D, E parâmetros de resultado nomeados

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
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
