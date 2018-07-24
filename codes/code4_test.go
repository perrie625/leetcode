package codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func combine(nums1 []int, nums2 []int) []int {
	var i, j int
	result := make([]int, 0, len(nums1)+len(nums2))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i = i + 1
		} else {
			result = append(result, nums2[j])
			j = j + 1
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i = i + 1
	}
	for j < len(nums2) {
		result = append(result, nums2[j])
		j = j + 1
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalNums := combine(nums1, nums2)
	totalLength := len(nums1) + len(nums2)
	mid := int(totalLength / 2)
	if totalLength%2 == 1 {
		return float64(totalNums[mid])
	}
	result := float64(float64(totalNums[mid]+totalNums[mid-1]) / 2.0)
	return result
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
