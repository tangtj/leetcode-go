package leetcode

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	i := len(nums) / 2
	return nums[i]
}
