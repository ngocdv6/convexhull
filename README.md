# Convex Hull using Graham's Scan

This package implements [Graham's scan](https://en.wikipedia.org/wiki/Graham_scan) in order to find the [convex hull](https://en.wikipedia.org/wiki/Convex_hull) of a given set of points in a Euclidean plane.

## Usage

Points containing an X and Y coordinate, along with a slice type for Points, are defined as follows:

```go
type Point struct {
	X, Y float64
}

type PointSlice []Point
```

Graham's scan takes a PointSlice and returns a PointSlice containing the Points contained by the convex hull. Example:

```go
MyConvexHull := convexhull.GrahamScan(MyPointSlice)
```