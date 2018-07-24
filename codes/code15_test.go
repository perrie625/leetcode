package codes

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	var left, right int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = len(nums) - 1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s > 0 {
				right -= 1
			} else if s < 0 {
				left += 1
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			}
		}
	}

	return result
}

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4}),
	)

	assert.Equal(
		[][]int{[]int{-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 0}),
	)

	assert.Equal(
		[][]int{[]int{0, 0, 0}},
		threeSum([]int{0, 0, 0, 0}),
	)
}
