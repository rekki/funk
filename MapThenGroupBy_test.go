package funk

import (
	"reflect"
	"strings"
	"testing"
)

//given a slice of words we map them to their length, transform to upper case and then group them by length
// e.g. one, two, three, four -> 3 -> [ONE, TWO], 4 -> [FOUR] 5 -> [THREE]

func TestMapThenGroupBy(t *testing.T) {
	type args struct {
		from   []string
		getKey func(s string) int
		mapFn  func(s string) string
	}
	var tests = []struct {
		name string
		args args
		want map[int][]string
	}{
		{name: "mapping and grouping numbers",
			args: args{
				from:   []string{"one", "two", "three", "four"},
				getKey: func(s string) int { return len(s) },
				mapFn:  func(s string) string { return strings.ToUpper(s) },
			},
			want: map[int][]string{4: {"FOUR"}, 3: {"ONE", "TWO"}, 5: {"THREE"}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapThenGroupBy(tt.args.from, tt.args.getKey, tt.args.mapFn); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("MapThenGroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
