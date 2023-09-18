package convexhull_test

import (
	"reflect"
	"testing"

	"github.com/jamesallured/convexhull"
)

func TestGrahamScan(t *testing.T) {
	tests := []struct {
		points, want []convexhull.Point
	}{
		{
			[]convexhull.Point{},
			nil,
		},
		{
			[]convexhull.Point{{0, 1}, {1, 0}, {0, 1}},
			[]convexhull.Point{{0, 1}, {1, 0}, {0, 1}},
		},
		{
			[]convexhull.Point{{1, 7}, {2, 3}, {4, 9}, {5, 2}, {5, 5}, {6, 7}, {7, 3}, {8, 8}, {10, 6}},
			[]convexhull.Point{{5, 2}, {7, 3}, {10, 6}, {8, 8}, {4, 9}, {1, 7}, {2, 3}},
		},
		{
			[]convexhull.Point{{3, 1}, {4, 1}, {4, 3}, {1, 2}, {4, 2}, {5, 2}, {2, 3}, {2, 2}, {3, 2}, {2, 1}, {3, 3}},
			[]convexhull.Point{{2, 1}, {4, 1}, {5, 2}, {4, 3}, {2, 3}, {1, 2}},
		},
	}
	for _, test := range tests {
		got := convexhull.GrahamScan(test.points)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}
