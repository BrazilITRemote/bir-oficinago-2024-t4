package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Matias")  // Não tenho o codigo ainda e ja estou pensando no futuro
		want := "Hello, Matias" //O que eu quero que aconteça no futuro

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
