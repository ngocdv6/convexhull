package convexhull

import (
	"testing"
)

func TestRotation(t *testing.T) {
	tests := []struct {
		x, y, z Point
		want    float64
	}{
		{Point{0, 0}, Point{1, 1}, Point{2, 2}, 0},
		{Point{0, 0}, Point{1, 1}, Point{2, 1}, -1},
		{Point{0, 0}, Point{1, 1}, Point{1, 2}, 1},
	}

	for _, test := range tests {
		got := rotation(test.x, test.y, test.z)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}

func TestSquareDistance(t *testing.T) {
	tests := []struct {
		x, y Point
		want float64
	}{
		{Point{0, 0}, Point{1, 1}, 2},
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 0}, Point{-1, -1}, 2},
	}

	for _, test := range tests {
		got := squareDistance(test.x, test.y)
		if got != test.want {
			t.Errorf("got %v, wanted %v", got, test.want)
		}
	}
}
