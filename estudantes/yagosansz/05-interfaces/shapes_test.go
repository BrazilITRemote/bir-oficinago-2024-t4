package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// Advantages of using a map instead of slice of structs to store test cases:
	//   1. "name" is simply the map index
	//   2. map iteration order is random, ensure tests are independent
	//   from one another!
	areaTests := map[string]struct {
		shape   Shape
		hasArea float64
	}{
		"Rectangle": {
			shape: Rectangle{12, 6}, hasArea: 72.0,
		},
		"Circle": {
			shape: Circle{10}, hasArea: 314.1592653589793,
		},
		"Triangle": {
			shape: Triangle{12, 6}, hasArea: 36.0,
		},
	}

	// Run specific test within the table with: "go test -run TestArea/Rectangle"
	// or "go test -v -run TestArea/Rectangle"
	for name, test := range areaTests {
		t.Run(name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.hasArea {
				t.Errorf("%#v got %g want %g", test.shape, got, test.hasArea)
			}
		})
	}
}
