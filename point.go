package convexhull

// Point represents a point in 2D space, with an X and Y co-ordinate each
// represented by a float64.
type Point struct {
	X, Y float64
}

// pointSlice implements sort.Interface.
// https://golang.org/pkg/sort/#Interface
type pointSlice []Point

// Len returns the number of Points in the PointSlice.
func (p pointSlice) Len() int {
	return len(p)
}

// Less returns true if the polar angle between the minimum Point in the slice
// and the Point at index i is less than the polar angle between the minimum
// Point in the slice and the Point at index j, otherwise it returns false. If
// the polar angles are the same, i.e. the two Points are collinear with the
// minumum Point, then Less returns true if the square distance between the
// minimum Point and the Point at index i is less than the square distance
// between the minumum Point and the Point at index j. It is assumed that the
// Point at index 0 is the minimum Point.
func (p pointSlice) Less(i, j int) bool {
	switch rotation(p[0], p[i], p[j]) {
	case 0:
		return squareDistance(p[0], p[i]) < squareDistance(p[0], p[j])
	default:
		return (p[i].X-p[0].X)*(p[j].Y-p[0].Y)-(p[j].X-p[0].X)*(p[i].Y-p[0].Y) > 0
	}
}

// Swap the point at index i with the point at index j.
func (p pointSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
