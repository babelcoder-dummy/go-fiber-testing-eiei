package demo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 1+1, "they should be equal")
	assert.NotEqual(t, 4, 1+2)
	assert.True(t, 1 == 0+1)
	assert.False(t, false)
	assert.Greater(t, 10, 2)
	assert.GreaterOrEqual(t, 5, 1)
	assert.Less(t, 3, 5)
	assert.LessOrEqual(t, 1, 2)

	var s1 []int
	assert.Nil(t, s1)

	s2 := []int{1, 2, 3}
	assert.NotNil(t, s2)
	assert.Contains(t, s2, 2)
	assert.Contains(t, "Hello World", "ll")
	assert.Len(t, s2, 3)
	assert.Subset(t, s2, []int{2, 1})

	err := errors.New("my error")
	assert.Error(t, err)
	assert.NoError(t, nil)
}
