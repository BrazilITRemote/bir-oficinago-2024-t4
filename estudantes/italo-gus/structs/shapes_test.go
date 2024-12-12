package structs // Interfaces

import (
	"testing"
)

/*
func TestPerimeter(t *testing.T) {
	received := Perimeter(10.0, 10.0)
	expected := 40.0
	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}

func TestArea(t *testing.T) {
	received := Area(12.0, 6.0)
	expected := 72.0

	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}
*/

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	//received := Perimeter(rectangle)
	received := rectangle.Perimeter()
	expected := 40.0
	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}

/*
func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	//received := Area(rectangle)
	received := rectangle.Area()
	expected := 72.0
	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}
*/

func TestArea(t *testing.T) {
	// função de teste utilizando a interface shape como parâmetro, declarada como variável
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		received := shape.Area()
		if received != expected {
			t.Errorf("expected '%g' but received '%g'", expected, received)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
