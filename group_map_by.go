package funk

// GroupMapBy groups mapped values by keys derived from the original values,
// preserving input order within each group.
func GroupMapBy[T any, K comparable, R any](
	items []T,
	key func(T) K,
	mapper func(T) R,
) map[K][]R {
	groups := make(map[K][]R)
	for _, item := range items {
		itemKey := key(item)
		groups[itemKey] = append(groups[itemKey], mapper(item))
	}
	return groups
}
