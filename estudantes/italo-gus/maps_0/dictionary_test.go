package maps_0 // Função de cadastro no mapa, aprimorada para verificar se a chave (índice) já é cadastrada

import (
	"testing"
)

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

/*
	func TestAdd(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	}
*/
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	received, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, received, definition)
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
