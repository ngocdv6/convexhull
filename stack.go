package convexhull

import "errors"

// stack represents a stack of Points which is equivalent to a slice of Points.
type stack []Point

// errEmptyStack is an error to be returned when an operation requiring an item
// to be on the stack is performed on an empty stack.
var errEmptyStack = errors.New("stack is empty")

// errSingleItemStack is an error to be retruend when an operation requiring
// more than a single item on the stack is performed on a stack containing only a
// single item
var errSingleItemStack = errors.New("stack only has one item")

// push a Point onto the top of the stack.
func (s *stack) push(p Point) {
	*s = append(*s, p)
}

// empty returns true if the stack is empty.
func (s *stack) empty() bool {
	return len(*s) == 0
}

// pop removes the Point at the top of the stack and returns it. If the stack is
// empty, then a zero-initialised Point will be returned along with an error.
func (s *stack) pop() (Point, error) {
	if s.empty() {
		return Point{}, errEmptyStack
	}
	index := len(*s) - 1
	point := (*s)[index]
	(*s) = (*s)[:index]
	return point, nil
}

// peek returns the Point at the top of the stack without removing it. If the
// stack is empty, then a zero-initialised Point will be returned along with an
// error.
func (s *stack) peek() (Point, error) {
	if s.empty() {
		return Point{}, errEmptyStack
	}
	index := len(*s) - 1
	point := (*s)[index]
	return point, nil
}

// peekBelow returns the Point below the top of the stack without removing it. If
// the stack has less than two elements, then a zero-initialised Point will be
// returned along with an error.
func (s *stack) peekBelow() (Point, error) {
	switch l := len(*s); {
	case l == 0:
		return Point{}, errEmptyStack
	case l == 1:
		return Point{}, errSingleItemStack
	}

	index := len(*s) - 2
	point := (*s)[index]
	return point, nil
}
