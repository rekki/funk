package funk

// GroupBy groups items by a derived key, preserving input order within each
// group. The key function is called exactly once per item.
func GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, item := range items {
		itemKey := key(item)
		groups[itemKey] = append(groups[itemKey], item)
	}
	return groups
}
