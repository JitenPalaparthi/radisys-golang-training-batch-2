package shape

type IShape interface {
	IArea
	IPerimeter
}

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}
