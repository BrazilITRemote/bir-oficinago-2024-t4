package arrays // Refactor - Refatoração / for range / _ blank identifier (identificador em branco)

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	repeated := Sum(numbers)
	expected := 15

	if repeated != expected {
		t.Errorf("expected '%d' but repeated '%d' given, '%v'", expected, repeated, numbers)
	}

}
