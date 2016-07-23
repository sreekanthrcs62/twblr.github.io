package interfaces

import "testing"

func TestAreaofSquare(t *testing.T) {
	s := Square{length: 2}
	result := s.Area()

	if result != 4 {
		t.Errorf("Expected area of square to be 4, got %d", result)
	}
}

func TestAreaofRectangle(t *testing.T) {
	r := Rectangle{length: 2, breadth: 10}
	result := r.Area()

	if result != 20 {
		t.Errorf("Expected area of rectangle to be 20, got %d", result)
	}
}

func TestAreaofComplexShape(t *testing.T) {
	s := Square{length: 10}
	r := Rectangle{length: 2, breadth: 3}

	h := Hybrid{&s, &r}
	result := h.Area()

	if result != 106 {
		t.Errorf("Expected area of hybrid to be 106, got %d", result)
	}
}
