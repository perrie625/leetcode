package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fillProcess(s string) string {
	result := make([]rune, 0, 2*len(s)+1)
	for _, c := range s {
		result = append(result, '#')
		result = append(result, c)
	}
	result = append(result, '#')
	return string(result)
}

func longestPalindromeManacher(s string) string {
	if len(s) <= 1 {
		return s
	}
	t := fillProcess(s)
	n := len(t)
	p := make([]int, n, n)
	var c, r = 0, 0
	var i int
	for i = 1; i < n-1; i++ {
		mirrorIndex := 2*c - i
		if r > i {
			p[i] = int(math.Min(float64(r-i), float64(p[mirrorIndex])))
		} else {
			p[i] = 0
		}

		// 继续往外判断是否可以延伸回文范围
		for i+1+p[i] < n && i-p[i]-1 >= 0 && t[i-p[i]-1] == t[i+1+p[i]] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	length := 0
	center := 0
	for i = 1; i < n; i++ {
		if p[i] > length {
			length = p[i]
			center = i
		}
	}

	return s[(center-length)/2 : (center-length)/2+length]
}

func TestManacher(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abba", longestPalindromeManacher("abba"))
	assert.Equal("a", longestPalindromeManacher("a"))
	assert.Equal("hereh", longestPalindromeManacher("ahdherehss"))
	assert.Equal("bab", longestPalindromeManacher("babad"))
}
