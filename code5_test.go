package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindrome(s string) string {
	runes := []rune(s)
	length := len(runes)
	if length < 2 {
		return s
	}
	var i, j, left, right int
	maxLen := 1
	maxStr := runes[:1]
	for i = 1; i < length && length-i > (maxLen/2); {
		left = i
		right = i
		// combine same rune
		for j = i - 1; j >= 0; j-- {
			if runes[j] != runes[i] {
				break
			}
			left--
		}
		for j = i + 1; j < length; j++ {
			if runes[j] != runes[i] {
				break
			}
			right++
		}

		for j = 1; left-j >= 0 && right+j < length; j++ {
			if runes[left-j] != runes[right+j] {
				break
			}
		}
		if len(runes[left-j+1:right+j]) > maxLen {
			maxStr = runes[left-j+1 : right+j]
			maxLen = len(maxStr)
		}
		i = right + 1
	}

	return string(maxStr)
}

func TestLongestPalindrome(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abba", longestPalindrome("abba"))
	assert.Equal("a", longestPalindrome("a"))
	assert.Equal("hereh", longestPalindrome("ahdherehss"))
	assert.Equal("bab", longestPalindrome("babad"))
}
