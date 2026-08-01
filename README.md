# funk

`funk` is a small, zero-dependency collection of type-safe helpers for Go slices,
maps, and sets.

> **Status:** `funk` follows semantic versioning. While the major version is
> `v0`, minor releases may include breaking API changes.

## Installation

```bash
go get github.com/rekki/funk
```

## Example

```go
package main

import (
	"fmt"

	"github.com/rekki/funk"
)

func main() {
	lengths := funk.Map([]string{"one", "two", "three"}, func(value string) int {
		return len(value)
	})

	byLength := funk.GroupBy([]string{"one", "two", "three"}, func(value string) int {
		return len(value)
	})

	fmt.Println(lengths)     // [3 3 5]
	fmt.Println(byLength[3]) // [one two]
}
```

## API

The package includes focused helpers for:

- transformation: `Map`, `Filter`, `FilterMap`, `FlatMap`, `Reduce`
- predicates and selection: `All`, `Best`, `Partition`, `Take`
- indexing and grouping: `IndexBy`, `Associate`, `GroupBy`, `GroupMapBy`
- set-like slice operations: `Unique`, `UniqueBy`, `Difference`, `Intersect`, `Union`, `SetEqual`, `DuplicateCounts`
- mutable sets: `Set`
- interoperability: `Box`

See the package documentation for complete behavior and examples.

## Semantics

- Functions that range over maps do not guarantee output order.
- Returned subslices may share storage with their input when documented.

## Contributing and security

See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidance and
[SECURITY.md](SECURITY.md) for reporting vulnerabilities.

## License

[MIT](LICENSE)
