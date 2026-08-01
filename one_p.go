package funk

// OneP returns the item selected by repeatedly applying replace. The replace
// function should report whether next should replace current. OneP panics when
// from is empty.
func OneP[A any](from []A, replace func(current, next A) bool) A {
	if len(from) == 0 {
		panic("funk: OneP called with an empty slice")
	}
	best := from[0]
	for _, next := range from[1:] {
		if replace(best, next) {
			best = next
		}
	}
	return best
}

// OneOrNil is the pointer-valued form of OneP. It returns nil when from is
// empty.
func OneOrNil[A any](from []*A, replace func(current, next *A) bool) *A {
	if len(from) == 0 {
		return nil
	}
	best := from[0]
	for _, next := range from[1:] {
		if replace(best, next) {
			best = next
		}
	}
	return best
}

// OneRef returns a reference to the first item in a slice if one exists, otherwise nil.
func OneRef[A any](from []*A) *A {
	if len(from) > 0 {
		return from[0]
	}
	return nil
}
