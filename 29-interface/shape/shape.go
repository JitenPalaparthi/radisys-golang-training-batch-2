package shape

import (
	"errors"
	"fmt"
)

func New(iShape IShape) (*Shape, error) {
	if iShape == nil {
		return nil, errors.New("nil iShape")
	}
	return &Shape{iShape: iShape}, nil
}

type Shape struct {
	iShape IShape // unexported filed ; becasue it stats with lowerCase
}

func (s *Shape) Area() {
	fmt.Println("Area:", s.iShape.Area())
}

func (s *Shape) Perimeter() {
	fmt.Println("Perimeter:", s.iShape.Perimeter())
}
