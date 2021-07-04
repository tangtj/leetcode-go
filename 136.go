package leetcode

func singleNumber(nums []int) int {
	t := nums[0]
	for l, j := 1, len(nums); l < j; l++ {
		t ^= nums[l]
	}
	return t
}
