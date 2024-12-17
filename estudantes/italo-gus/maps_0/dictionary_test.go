package maps_0

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

	/*
		t.Run("unknown word", func(t *testing.T) {
			_, err := dictionary.Search("unknown")
			expected := "could not find the word you were looking for"

			if err == nil {
				t.Fatal("expected to get an error.")
			}

			assertStrings(t, err.Error(), expected)
		})
	*/

	t.Run("unknown word", func(t *testing.T) {
		_, received := dictionary.Search("unknown")
		if received == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, received, ErrNotFound)
		// gerando erro, criando erro diferente do esperado
		//assertError(t, received, errors.New("could not find the word you were looking for !"))
	})

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
