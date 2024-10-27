package hello_world

import "testing"

func TestOla(t *testing.T) {
	result := Hello("Alexandre")
	expected := "Hello Alexandre"

	if result != expected {
		t.Errorf("resultado %q, esperado %q", result, expected)
	}
}
