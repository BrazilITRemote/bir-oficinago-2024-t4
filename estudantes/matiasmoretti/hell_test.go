package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Matias")  // Não tenho o codigo ainda e ja estou pensando no futuro
	want := "Hello, Matias" //O que eu quero que aconteça no futuro

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
