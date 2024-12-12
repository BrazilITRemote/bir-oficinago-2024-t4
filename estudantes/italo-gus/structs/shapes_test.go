package structs //
import "testing"

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
	received := Perimeter(rectangle)
	expected := 40.0
	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	received := Area(rectangle)
	expected := 72.0
	if received != expected {
		t.Errorf("expected '%.2f' but received '%.2f'", expected, received)
	}
}
