// Given a list of non negative integers, arrange them such that they form the largest number.

// Example 1:

// Input: [10,2]
// Output: "210"
// Example 2:

// Input: [3,30,34,5,9]
// Output: "9534330"
// Note: The result may be very large, so you need to return a string instead of an integer.

package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}
	sort.Slice(strNums, func(i, j int) bool { return strNums[i]+strNums[j] > strNums[j]+strNums[i] })
	result := strings.Join(strNums, "")
	if result[0] == '0' {
		result = "0"
	}
	return result
}

func TestLargestNumber(t *testing.T) {
	assert.Equal(t, "210", largestNumber([]int{10, 2}))
	assert.Equal(t, "4540", largestNumber([]int{40, 45}))
	assert.Equal(t, "0", largestNumber([]int{0, 0}))
	assert.Equal(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}
