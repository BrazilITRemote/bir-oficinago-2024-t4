package maps_0

import (
	"testing"
)

/*
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	// meuMapa :=  map [tipo_de_dados_da_chave(índice)] tipo_de_dados_do_elemento
	// https://go.dev/ref/spec#Map_types
	// https://pkg.go.dev/maps
	// https://go.dev/doc/effective_go#maps
	// https://go.dev/blog/maps
	//
	received := Search(dictionary, "test")
	expected := "this is just a test"

	if received != expected {
		t.Errorf("expected %q but received %q given %q ", expected, received, "test")
	}
}
*/

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	received := Search(dictionary, "test")
	expected := "this is just a test"

	assertStrings(t, received, expected)
}

func assertStrings(t testing.TB, received, expected string) {
	t.Helper()

	if received != expected {
		t.Errorf("expected %q but received %q ", expected, received)
	}
}
