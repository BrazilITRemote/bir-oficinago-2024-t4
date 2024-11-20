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

func TestRepeatCount(t *testing.T) {
	repeated := RepeatCount("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeatCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestExampleRepeat(t *testing.T) {
	repeated := ExampleRepeat()
	expected := true

	if repeated != expected {
		t.Errorf("expected '%t' but got '%t'", expected, repeated)
	}
}

/*
func BenchmarkExampleRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleRepeat()
	}
}
*/
