package leetcode

func minStartValue(nums []int) int {

	startValue := 0

	sum := 0
	for i, l := 0, len(nums); i < l; i++ {
		sum += nums[i]
		if sum <= 0 {
			startValue += 1 - sum
			sum = 1
		}
	}
	if startValue == 0 {
		return 1
	}
	return startValue
}
