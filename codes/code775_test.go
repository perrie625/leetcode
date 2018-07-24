package codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIdealPermutation(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	for i, num := range A {
		if num-i > 1 || num-i < -1 {
			return false
		}
	}
	return true
}

func Test775(t *testing.T) {
	assert.True(t, isIdealPermutation([]int{0}))
	assert.True(t, isIdealPermutation([]int{1, 0, 2}))
	assert.True(t, isIdealPermutation([]int{0, 2, 1, 3, 4}))
	assert.False(t, isIdealPermutation([]int{5, 0, 2, 1, 3, 4}))
}
