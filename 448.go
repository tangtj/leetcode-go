package leetcode

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, len(nums))
	for _, num := range nums {
		ret[num] = ret[num] + 1
	}
	r := make([]int, 0)
	for i, l := 0, len(nums); i < l; i++ {
		if nums[i] == 0 {
			r = append(r, i)
		}
	}
	return r
}
