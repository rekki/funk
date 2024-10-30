package funk

type Set[T comparable] struct {
	items map[T]struct{}
}

func (s *Set[T]) Values() []T {
	list := make([]T, 0, len(s.items))
	for v := range s.items {
		list = append(list, v)
	}
	return list
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.items, v)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) AddMulti(items ...T) {
	for _, v := range items {
		s.Add(v)
	}
}

type FilterFunc[T comparable] func(v T) bool

func (s *Set[T]) Filter(P FilterFunc[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.items {
		if !P(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.items {
		res.Add(v)
	}
	for v := range s2.items {
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.items {
		if !s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.items {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func NewSet[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{}
	s.items = make(map[T]struct{})
	if initial != nil {
		s.AddMulti(initial...)
	}
	return s
}
