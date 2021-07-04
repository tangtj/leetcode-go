package leetcode

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	maxSum := nums[0]
	sum := nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		sum = MathMax(nums[i], sum+nums[i])
		maxSum = MathMax(sum, maxSum)
	}
	return maxSum
}
