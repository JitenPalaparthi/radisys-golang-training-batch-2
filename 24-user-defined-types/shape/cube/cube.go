package cube

import "errors"

func New(l, b, h float32) (*Cube, error) {

	if l == 0 || b == 0 || h == 0 {
		return nil, errors.New("l,b or h cannot be zero")
	}

	return &Cube{L: l, B: b, H: h}, nil
}

type Cube struct {
	L, B, H float32
	A, P    float32
}

func (c *Cube) Area() float32 {
	c.A = c.L * c.B * c.H
	return c.A
}

func (c *Cube) Perimeter() float32 {
	c.P = 4 * (c.L + c.B + c.H)
	return c.P
}

func NewCuboid() (*Cuboid, error) {
	return nil, nil
}

type Cuboid struct{}
