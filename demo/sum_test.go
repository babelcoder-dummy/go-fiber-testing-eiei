package demo_test

import (
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/demo"
)

func TestSum(t *testing.T) {
	r := demo.Sum([]uint{1, 2, 3})

	if r != 6 {
		t.Errorf("got %d; expected %d", r, 6)
	}
}
