package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		t := target - num
		if idx, ok := m[t]; ok && idx != i {
			return []int{i, idx}
		}
	}
	return []int{0, 0}
}
