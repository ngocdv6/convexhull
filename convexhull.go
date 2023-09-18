package convexhull

import (
	"log"
	"sort"
)

// GrahamScan returns a PointSlice containing the Points belonging to the
// convex hull of a finite set of Points. The Graham scan uses a stack to
// remove concave points from the hull.
//
// See: https://en.wikipedia.org/wiki/Graham_scan
func GrahamScan(points []Point) []Point {
	ps := pointSlice(points)
	// If there are less than three Points, it's impossible to find a convex hull
	// If there are three Points, the input is the convex hull
	switch l := len(points); {
	case l < 3:
		return nil
	case l == 3:
		return ps
	}

	// Find the lowest, leftmost Point
	for i, point := range ps {
		if point.Y < ps[0].Y || point.Y == ps[0].Y && point.X < ps[0].X {
			ps.Swap(0, i)
		}
	}

	// Sort according to the polar angle relative to points[0]
	sort.Sort(ps)

	// Remove collinear Points - keep the furthest Point
	for i := 1; i < len(ps)-1; {
		switch {
		case rotation(ps[0], ps[i], ps[i+1]) == 0:
			ps = append(ps[:i], ps[i+1:]...)
		default:
			i++
		}
	}

	var s stack

	// Compute the convex hull by removing concavities from the set of Points
	for _, point := range ps {
		for len(s) > 1 {
			below, err := s.peekBelow()
			if err != nil {
				log.Fatal(err)
			}
			top, err := s.peek()
			if err != nil {
				log.Fatal(err)
			}
			if rotation(below, top, point) <= 0 {
				_, _ = s.pop()
			} else {
				break
			}
		}
		s.push(point)
	}

	return pointSlice(s)
}

// rotation returns the orientation of the direction of three points. If the
// points are collinear, then 0 is returned. If the points turn clockwise, a
// value less than 0 is returned. If the points turn counterclockwise, a value
// greater than 0 is returned
func rotation(x, y, z Point) float64 {
	return (y.X-x.X)*(z.Y-x.Y) - (z.X-x.X)*(y.Y-x.Y)
}

// squareDistance returns the Euclidean squared distance between two points.
func squareDistance(x, y Point) float64 {
	dx, dy := x.X-y.X, x.Y-y.Y
	return (dx * dx) + (dy * dy)
}
