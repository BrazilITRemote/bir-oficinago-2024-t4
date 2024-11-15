package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating 'a' 5 times: 'aaaaa'", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeating 'a' 0 times: '' ", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeating '' 3 times: ''", func(t *testing.T) {
		repeated := Repeat("", 3)
		expected := ""

		assertCorrectMessage(t, repeated, expected)
	})

}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	characters := Repeat("c", 3)
	fmt.Println(characters)
	// Output: ccc
}
