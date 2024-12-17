package maps_0 // Evitar erro do tipo pânico em tempo de execução devido tipo dados mapa nula (nil)

import (
	"testing"
)

/*
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	received := dictionary.Search("test")
	expected := "this is just a test"

	assertStrings(t, received, expected)
}
*/

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		received, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, received, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, received := dictionary.Search("unknown")
		if received == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, received, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	// Nunca deve inicializar uma variável do tipo dados mapa nula (nil),
	// pois tipo de dadosmapa nil causarão um erro do tipo pânico em tempo de execução.
	// conforme https://go.dev/blog/maps devido a
	// "Um valor de mapa é um ponteiro para uma estrutura runtime.hmap."
	// citado em https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
	//
	// var dictionary map[string]string           (forma errada)
	//
	// Em vez disso, deve-se inicializar um mapa vazio {}
	// var dictionary = map[string]string{}	      (forma correta 1)
	//
	// Ou usando a palavra-chave make para criar um mapa
	// var dictionary = make(map[string]string)   (forma correta 2)

	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	received, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, received, expected)
}

func assertStrings(t testing.TB, received, expected string) {
	t.Helper()

	if received != expected {
		t.Errorf("expected %q but received %q ", expected, received)
	}
}

func assertError(t testing.TB, received, expected error) {
	t.Helper()

	if received != expected {
		t.Errorf("got error %q expected %q", received, expected)
	}
}
