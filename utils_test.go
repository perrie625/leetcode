package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, BinarySearch([]int{1, 2, 3}, 3))
}
