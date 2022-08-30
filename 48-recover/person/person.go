package person

// panicing inside a package is not a good practice..
// instead of panic , should return as an error
// let the caller handle it in which ever the way it has to be handled according to
// the requirement
func FullName(fn *string, ln *string) *string {
	if fn == nil {
		panic("firstName is nil")
	}
	if ln == nil {
		panic("lastName is nil")
	}

	fullname := *fn + *ln
	return &fullname
}
