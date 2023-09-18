package convexhull

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		s    stack
		in   Point
		want stack
	}{
		{stack{}, Point{0, 0}, stack{Point{0, 0}}},
		{stack{Point{0, 0}}, Point{1, 1}, stack{Point{0, 0}, Point{1, 1}}},
	}

	for _, test := range tests {
		test.s.push(test.in)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("wanted %v, got %v", test.want, test.s)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		stack stack
		want  bool
	}{
		{stack{}, true},
		{stack{Point{0, 0}}, false},
	}

	for _, test := range tests {
		got := test.stack.empty()
		if got != test.want {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		stack stack
		want  Point
		err   error
		after stack
	}{
		{stack{Point{0, 0}}, Point{0, 0}, nil, stack{}},
		{stack{Point{0, 0}, Point{1, 1}}, Point{1, 1}, nil, stack{Point{0, 0}}},
		{stack{}, Point{}, errEmptyStack, stack{}},
	}

	for _, test := range tests {
		got, err := test.stack.pop()
		if err != test.err {
			t.Errorf("wanted %v err, got %v err", test.err, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
		if !reflect.DeepEqual(test.stack, test.after) {
			t.Errorf("wanted %v, got %v after", test.after, test.stack)
		}
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		stack stack
		want  Point
		err   error
		after stack
	}{
		{stack{Point{0, 0}}, Point{0, 0}, nil, stack{Point{0, 0}}},
		{stack{Point{0, 0}, Point{1, 1}}, Point{1, 1}, nil, stack{Point{0, 0}, Point{1, 1}}},
		{stack{}, Point{}, errEmptyStack, stack{}},
	}

	for _, test := range tests {
		got, err := test.stack.peek()
		if err != test.err {
			t.Errorf("wanted %v err, got %v err", test.err, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
		if !reflect.DeepEqual(test.stack, test.stack) {
			t.Errorf("wanted %v, got %v after", test.after, test.stack)
		}
	}
}

func TestPeekBelow(t *testing.T) {
	tests := []struct {
		stack stack
		want  Point
		err   error
		after stack
	}{
		{stack{Point{0, 0}, Point{1, 1}}, Point{0, 0}, nil, stack{Point{0, 0}, Point{1, 1}}},
		{stack{Point{0, 0}}, Point{}, errSingleItemStack, stack{Point{0, 0}}},
		{stack{}, Point{}, errEmptyStack, stack{}},
	}

	for _, test := range tests {
		got, err := test.stack.peekBelow()
		if err != test.err {
			t.Errorf("wanted %v err, got %v err", test.err, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
		if !reflect.DeepEqual(test.stack, test.stack) {
			t.Errorf("wanted %v, got %v after", test.after, test.stack)
		}
	}
}
