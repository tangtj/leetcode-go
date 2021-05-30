package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	num := make([]int, 0)
	for x != 0 {
		num = append(num, x%10)
		x = x / 10
	}
	l := len(num)
	for i, j := 0, l-1; i <= j; {
		if num[i] != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}
