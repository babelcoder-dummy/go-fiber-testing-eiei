package demo_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
	suite.Suite
}

func (s *MyTestSuite) SetupSuite() {
	s.T().Log("SetupSuite")
}

func (s *MyTestSuite) SetupTest() {
	s.T().Log("SetupTest")
}

func (s *MyTestSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}

func (s *MyTestSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

func (s *MyTestSuite) TestOne() {
	s.T().Log("TestOne")
	s.Equal(1, 1)
}

func (s *MyTestSuite) TestTwo() {
	s.T().Log("TestTwo")
	s.Equal(2, 2)
}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
