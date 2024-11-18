package integers // Inteiros, Formatação de erros

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
		/* %q imprimir strings, %d imprimir inteiro.
		Formatação de erros
		https://pkg.go.dev/fmt#hdr-Printing
		*/
	}
}
