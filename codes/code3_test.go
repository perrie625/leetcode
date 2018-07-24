package codes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func inString(runes []rune, r rune) int {
	for i, c := range runes {
		if c == r {
			return i
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	var result, head, tail int
	for i, c := range runes {
		tail = i
		index := inString(runes[head:tail], c)
		if index >= 0 {
			head = head + index + 1
		}
		if (tail - head + 1) > result {
			result = tail - head + 1
		}
	}
	return result
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(lengthOfLongestSubstring("aav"), 2)
	assert.Equal(lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(lengthOfLongestSubstring("abcabcbb"), 3)
}
