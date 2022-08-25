package square

type Square float32

func (s *Square) Area() float32 {
	return float32(*s) * float32(*s)
}

func (s *Square) Perimeter() float32 {
	return 4 * float32(*s)
}
