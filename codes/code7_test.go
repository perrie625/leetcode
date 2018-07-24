package codes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(x int) int {
	if x > 2147483647 || x < -2147483648 {
		return 0
	}
	if x < 10 && x > -10 {
		return x
	}
	var minues bool
	if x < 0 {
		x = -x
		minues = true
	}

	digits := make([]int, 0)
	for x > 0 {
		digits = append(digits, x%10)
		x = x / 10
	}

	result := 0
	for i := 0; i < len(digits); i++ {
		result = result + digits[i]*int(math.Pow10(len(digits)-i-1))
	}
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	if minues {
		return -result
	}

	return result
}
func TestReverse(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, reverse(1))
	assert.Equal(134, reverse(431))
	assert.Equal(-134, reverse(-431))
	assert.Equal(0, reverse(-429383294389531))
}

// import "math"

// func reverse(x int) int {
// 	var flag int64
// 	if x > 0 {
// 		flag = 1
// 	} else if x < 0 {
// 		flag = -1
// 	}
// 	tmp := flag * int64(x)
// 	var ret int64
// 	// 将所有的负数的情况统一转换处理
// 	for tmp > 0 {
// 		ret = ret*10 + tmp%10
// 		tmp = (tmp - tmp%10) / 10
// 		if ret > math.MaxInt32 || ret*flag < math.MinInt32 {
// 			return 0
// 		}
// 	}
// 	return int(ret * flag)
// }
