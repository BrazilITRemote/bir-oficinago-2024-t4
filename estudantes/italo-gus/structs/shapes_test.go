package structs // TableDrivenTests

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

/*
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

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{12.0, 6.0}
		checkArea(t, triangle, 36)
	})

}
*/

/*
func TestArea(t *testing.T) {
	// TableDrivenTests - aplicando Testes orientados por tabela
	// https://go.dev/wiki/TableDrivenTests

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		received := tt.shape.Area()
		//fmt.Printf("BIR - type of TT: %s\n", reflect.TypeOf(tt.shape)) // Verificação de tipo
		//fmt.Printf("BIR - Testando: %T \n", tt.shape) // Impressão de tipo
		if received != tt.expected {
			t.Errorf("expected '%g' but received '%g'", tt.expected, received)
		}
	}

}
*/

func TestArea(t *testing.T) {
	// TableDrivenTests - aplicando Testes orientados por tabela
	// e cada teste com nome
	// https://go.dev/wiki/TableDrivenTests
	// https://pkg.go.dev/testing#T.Run

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		// usando tt.name de cada caso para usá-lo como o `t.Run`.nome_do_teste
		t.Run(tt.name, func(t *testing.T) {
			received := tt.shape.Area()
			if received != tt.hasArea {
				t.Errorf("%#v received %g expected %g", tt.shape, received, tt.hasArea)
			}
		})

	}
}
