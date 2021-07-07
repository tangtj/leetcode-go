package leetcode

func moveZeroes(nums []int) {
	l, r, lens := 0, 0, len(nums)
	for r < lens {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
