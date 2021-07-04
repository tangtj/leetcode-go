package leetcode

func search(nums []int, target int) int {
	l := len(nums)
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, l-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
