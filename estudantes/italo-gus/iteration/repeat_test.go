package iteration // Benchmarks em Go

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/* benchmarks em Go
https://pkg.go.dev/testing#hdr-Benchmarks
Para executar os benchmarks go test -bench=.
(no Windows Powershell go test -bench=".")
*/
