package funk

// Box converts xs to a slice of any values, as required by APIs such as
// database/sql.
func Box[T any](xs ...T) []any {
	return Map(xs, func(x T) any { return x })
}
