package funk

// Set is a mutable collection of distinct comparable values. Its zero value is
// ready to use. A Set must not be copied after first use because copies share
// the same backing map. Set is not safe for concurrent mutation. Method
// receivers and Set operands must be non-nil.
type Set[T comparable] struct {
	items map[T]struct{}
}

// Values returns the values in s in unspecified order.
func (s *Set[T]) Values() []T {
	list := make([]T, 0, len(s.items))
	for v := range s.items {
		list = append(list, v)
	}
	return list
}

// Has reports whether v is in s.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

// Add inserts v into s.
func (s *Set[T]) Add(v T) {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
	s.items[v] = struct{}{}
}

// Remove deletes v from s.
func (s *Set[T]) Remove(v T) {
	delete(s.items, v)
}

// Clear removes all values from s.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Size returns the number of values in s.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// AddMulti inserts every value in items into s.
func (s *Set[T]) AddMulti(items ...T) {
	for _, v := range items {
		s.Add(v)
	}
}

// FilterFunc reports whether a value should be retained by Set.Filter.
type FilterFunc[T any] func(v T) bool

// Filter returns a new set containing values accepted by predicate.
func (s *Set[T]) Filter(predicate FilterFunc[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.items {
		if !predicate(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

// Union returns a new set containing values from s and s2.
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

// Intersect returns a new set containing values present in both s and s2.
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

// Difference returns a new set containing values in s but not s2.
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

// NewSet returns a set initialized with initial.
func NewSet[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{}, len(initial))}
	s.AddMulti(initial...)
	return s
}
