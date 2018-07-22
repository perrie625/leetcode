package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	// change
	if len1 > len2 {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		ltmp := len1
		len1 = len2
		len2 = ltmp
	}
	imin := 0
	imax := len1
	middle := (len1 + len2 + 1) / 2
	var i, j, leftMax, rightMin int
	for imin <= imax {
		i = (imin + imax) / 2
		j = middle - i
		if i > imin && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else if i < imax && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else {
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				if nums1[i-1] >= nums2[j-1] {
					leftMax = nums1[i-1]
				} else {
					leftMax = nums2[j-1]
				}
			}
			if (len1+len2)%2 == 1 {
				return float64(leftMax)
			}
			if i == len1 {
				rightMin = nums2[j]
			} else if j == len2 {
				rightMin = nums1[i]
			} else {
				if nums1[i] <= nums2[j] {
					rightMin = nums1[i]
				} else {
					rightMin = nums2[j]
				}
			}
			return float64(leftMax+rightMin) / 2.0
		}
	}

	return 0.0
}

func TestFindMedianSortedArrays2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2.5, findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
	assert.Equal(3.0, findMedianSortedArrays2([]int{1}, []int{3, 4}))
	assert.Equal(4.0, findMedianSortedArrays2([]int{1, 5, 5}, []int{3, 4}))
	assert.Equal(4.5, findMedianSortedArrays2([]int{1, 5, 5, 5}, []int{3, 4}))
	assert.Equal(5.0, findMedianSortedArrays2([]int{1, 5, 5, 5, 5}, []int{3, 4}))
	assert.Equal(4.5, findMedianSortedArrays2([]int{1, 6, 7, 8}, []int{2, 3, 4, 5}))
}
