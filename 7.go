package leetcode

import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	ret := 0

	const max = math.MaxInt32 / 10
	const min = math.MinInt32 / 10

	for x != 0 {
		t := x % 10
		if ret > max || (ret == max && t > 7) {
			return 0
		} else if ret < min || (ret == min && t < -8) {
			return 0
		}
		ret = ret*10 + t
		x /= 10
	}
	return ret
}
