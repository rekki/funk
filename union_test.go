package funk

import (
	"reflect"
	"testing"
)

type Thing struct {
	Name string
	ID   int
}

func TestUnion(t *testing.T) {
	type args struct {
		xs []string
		ys []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Union works as expected with non empty params",
			args{[]string{"A", "B", "C"}, []string{"B", "C", "D", "E"}},
			[]string{"A", "B", "C", "D", "E"},
		},
		{
			"Union returns slice with xs values when ys is nil",
			args{[]string{"A", "B", "C"}, nil},
			[]string{"A", "B", "C"},
		},
		{
			"Union returns slice with ys values when xs is nil",
			args{nil, []string{"C", "D", "E"}},
			[]string{"C", "D", "E"},
		},
		{
			"Union returns empty slice if both are nil",
			args{nil, nil},
			[]string{},
		},
		{
			"Union returns empty slice if both are empty",
			args{[]string{}, []string{}},
			[]string{},
		},
		{
			"Union removes duplicates within each input",
			args{[]string{"A", "A", "B"}, []string{"B", "C", "C"}},
			[]string{"A", "B", "C"},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Union(tt.args.xs, tt.args.ys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
