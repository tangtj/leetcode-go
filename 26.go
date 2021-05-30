package leetcode

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	last := nums[0]
	j := len(nums)
	for i := 1; i < j; {
		if last == nums[i] {
			nums = append(nums[0:i], nums[i+1:j]...)
			j--
		} else {
			last = nums[i]
			i++
		}
	}
	return len(nums)
}
