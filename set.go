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
	values := make([]T, 0, len(s.items))
	for value := range s.items {
		values = append(values, value)
	}
	return values
}

// Contains reports whether value is in s.
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.items[value]
	return exists
}

// Add inserts value into s.
func (s *Set[T]) Add(value T) {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
	s.items[value] = struct{}{}
}

// Remove deletes value from s.
func (s *Set[T]) Remove(value T) {
	delete(s.items, value)
}

// Clear removes all values from s.
func (s *Set[T]) Clear() {
	clear(s.items)
}

// Len returns the number of values in s.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// AddAll inserts every value in items into s.
func (s *Set[T]) AddAll(items ...T) {
	if s.items == nil {
		s.items = make(map[T]struct{}, len(items))
	}
	for _, value := range items {
		s.items[value] = struct{}{}
	}
}

// Filter returns a new set containing values accepted by predicate.
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
	result := newSet[T](len(s.items))
	for value := range s.items {
		if predicate(value) {
			result.items[value] = struct{}{}
		}
	}
	return result
}

// Union returns a new set containing values from s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := newSet[T](len(s.items) + len(other.items))
	for value := range s.items {
		result.items[value] = struct{}{}
	}
	for value := range other.items {
		result.items[value] = struct{}{}
	}
	return result
}

// Intersect returns a new set containing values present in both s and other.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	smaller, larger := s, other
	if len(smaller.items) > len(larger.items) {
		smaller, larger = larger, smaller
	}

	result := newSet[T](len(smaller.items))
	for value := range smaller.items {
		if larger.Contains(value) {
			result.items[value] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set containing values in s but not other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := newSet[T](len(s.items))
	for value := range s.items {
		if !other.Contains(value) {
			result.items[value] = struct{}{}
		}
	}
	return result
}

// NewSet returns a set initialized with initial.
func NewSet[T comparable](initial ...T) *Set[T] {
	set := newSet[T](len(initial))
	set.AddAll(initial...)
	return set
}

func newSet[T comparable](capacity int) *Set[T] {
	return &Set[T]{items: make(map[T]struct{}, capacity)}
}
