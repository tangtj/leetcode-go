package leetcode

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if target <= nums[0] {
		return 0
	} else if target > nums[l-1] {
		return l
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
	return l
}
