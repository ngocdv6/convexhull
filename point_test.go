package convexhull

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		in   pointSlice
		want int
	}{
		{pointSlice{Point{0, 0}}, 1},
		{pointSlice{}, 0},
	}

	for _, test := range tests {
		got := test.in.Len()
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}

func TestLess(t *testing.T) {
	tests := []struct {
		in   pointSlice
		want bool
	}{
		{pointSlice{Point{0, 0}, Point{1, 1}, Point{2, 2}}, true},
		{pointSlice{Point{0, 0}, Point{2, 2}, Point{1, 1}}, false},
		{pointSlice{Point{0, 0}, Point{1, 1}, Point{1, 2}}, true},
		{pointSlice{Point{0, 0}, Point{1, 1}, Point{2, 1}}, false},
		{pointSlice{Point{0, 0}, Point{1, 1}, Point{1, 1}}, false},
	}

	for _, test := range tests {
		got := test.in.Less(1, 2)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		in   pointSlice
		want pointSlice
	}{
		{pointSlice{Point{0, 0}, Point{1, 1}}, pointSlice{Point{1, 1}, Point{0, 0}}},
	}

	for _, test := range tests {
		test.in.Swap(0, 1)
		if !reflect.DeepEqual(test.in, test.want) {
			t.Errorf("got %v, wanted %v", test.in, test.want)
		}
	}
}
