package interfaces

type Shape interface {
	Area() int
}

type Square struct {

	length int
	 
}

func (s *Square) Area() int {
	return s.length * s.length
}


type Rectangle struct {

	length int
	breadth int
	 
}


type Hybrid struct {
	// Better way is to add a slice of type shape list and parse that 
	s Shape
	r Shape
}

func (r *Rectangle) Area() int {
	return r.length * r.breadth
}

func (h *Hybrid) Area() int {
	return getArea(h.s) + getArea(h.r)
}


func getArea(s Shape) int {
	return s.Area()
}