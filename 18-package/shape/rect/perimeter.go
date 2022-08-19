package rect

func Perimeter(l, b float32) float32 {
	return 2 * (l + b)
}

func PerimeterGlobal() float32 {
	return 2 * (Length + Width)
}
