package main // tdd com argumento

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris") // variáve que recebe o retorno do resultado da função a ser testada, passando argumento para função
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
