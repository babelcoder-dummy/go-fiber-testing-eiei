package demo_test

import (
	"fmt"
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/demo"
	"github.com/stretchr/testify/require"
)

func TestCalculateFare(t *testing.T) {
	tests := []struct {
		kind     string
		n        uint
		expected uint
	}{
		{kind: "MRT", n: 1, expected: 14},
		{kind: "MRT", n: 5, expected: 22},
		{kind: "BTS", n: 1, expected: 18},
		{kind: "Airport Link", n: 7, expected: 34},
		{kind: "SRT", n: 10, expected: 0},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s %d stations", tc.kind, tc.n), func(t *testing.T) {
			t.Logf("%s %d stations", tc.kind, tc.n)
			got := demo.CalculateFare(tc.kind, tc.n)

			require.Equalf(t, tc.expected, got, "got %d; want %d", got, tc.expected)
		})
	}
}
