package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	slen := len(s)
	if slen <= numRows {
		return s
	}
	term := numRows*2 - 2
	var i, j, temp int
	result := make([]byte, 0, slen)
	for i = 0; i < numRows; i++ {
		for j = i; j < slen; j = j + term {
			result = append(result, s[j])
			temp = (numRows-i-1)*2 + j
			if i != 0 && i != numRows-1 && temp < slen {
				result = append(result, s[temp])
			}
		}
	}

	return string(result)
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abced", convert("abcde", 4))
	assert.Equal("acbd", convert("abcd", 2))
	assert.Equal("a", convert("a", 1))
	assert.Equal("ab", convert("ab", 1))
	assert.Equal("PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}
