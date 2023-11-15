package demo_test

import (
	"fmt"
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/demo"
	"github.com/stretchr/testify/suite"
)

type CircleTestSuite struct {
	suite.Suite
}

func (suite *CircleTestSuite) TestArea() {

	tests := []struct {
		raduis   float64
		expected string
	}{
		{raduis: 10, expected: "314.16"},
		{raduis: 7, expected: "153.94"},
	}

	for _, tc := range tests {
		circle := demo.Circle{Radius: tc.raduis}
		got := fmt.Sprintf("%.2f", circle.Area())

		suite.Equal(tc.expected, got)
	}
}

func (suite *CircleTestSuite) TestCircumference() {
	tests := []struct {
		raduis   float64
		expected string
	}{
		{raduis: 10, expected: "62.83"},
		{raduis: 7, expected: "43.98"},
	}

	for _, tc := range tests {
		circle := demo.Circle{Radius: tc.raduis}
		got := fmt.Sprintf("%.2f", circle.Circumference())

		suite.Equal(tc.expected, got)
	}
}

func TestCircleTestSuite(t *testing.T) {
	suite.Run(t, new(CircleTestSuite))
}
