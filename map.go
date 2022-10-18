package funk

func Map[A any, From any](from []From, mapper func(From) A) []A {
	out := []A{}

	for _, item := range from {
		v := mapper(item)
		out = append(out, v)
	}

	return out
}

func Unique[A comparable](items []A) []A {
	out := []A{}
	values := map[A]bool{}
	for _, item := range items {
		_, isDuplicate := values[item]
		if !isDuplicate {
			out = append(out, item)
			values[item] = true
		}
	}
	return out
}
