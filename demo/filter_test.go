package demo_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/demo"
)

type filterTest[T comparable] struct {
	items    []T
	fn       func(item T) bool
	expected []T
}

func TestFilter_uint(t *testing.T) {
	t.Parallel()
	tests := []filterTest[uint]{
		{
			items:    []uint{1, 2, 3},
			fn:       func(item uint) bool { return item > 1 },
			expected: []uint{2, 3},
		},
		{
			items:    []uint{1, 3, 5, 7, 9},
			fn:       func(item uint) bool { return item > 5 },
			expected: []uint{7, 9},
		},
		{
			items:    []uint{1, 2, 3, 4, 5, 7, 9},
			fn:       func(item uint) bool { return item%2 == 0 },
			expected: []uint{2, 4},
		},
	}

	for _, tc := range tests {
		got := demo.Filter(tc.items, tc.fn)

		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v; expected %v", got, tc.expected)
		}
	}

}

func TestFilter_string(t *testing.T) {
	t.Parallel()
	tests := []filterTest[string]{
		{
			items:    []string{"hello", "world"},
			fn:       func(item string) bool { return strings.Contains(item, "ll") },
			expected: []string{"hello"},
		},
		{
			items:    []string{"My", "Name", "Is", "Babel", "Coder"},
			fn:       func(item string) bool { return len(item) > 3 },
			expected: []string{"Name", "Babel", "Coder"},
		},
	}

	for _, tc := range tests {
		got := demo.Filter(
			tc.items,
			tc.fn,
		)

		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v; expected %v", got, tc.expected)
		}
	}
}
