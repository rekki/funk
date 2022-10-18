package funk

func FlatMap[A any, From any](from []From, mapper func(From) []A) []A {
	out := []A{}
	for _, item := range from {
		out = append(out, mapper(item)...)
	}
	return out
}
