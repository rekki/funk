package funk

import "fmt"

func ExampleMap() {
	lengths := Map([]string{"one", "two", "three"}, func(value string) int {
		return len(value)
	})
	fmt.Println(lengths)
	// Output: [3 3 5]
}

func ExampleGroupBy() {
	groups := GroupBy([]string{"one", "two", "three", "four"}, func(value string) int {
		return len(value)
	})
	fmt.Println(groups[3])
	fmt.Println(groups[4])
	fmt.Println(groups[5])
	// Output:
	// [one two]
	// [four]
	// [three]
}

func ExampleAssociate() {
	lengths := Associate([]string{"one", "three"}, func(value string) (string, int) {
		return value, len(value)
	})
	fmt.Println(lengths["one"])
	fmt.Println(lengths["three"])
	// Output:
	// 3
	// 5
}
