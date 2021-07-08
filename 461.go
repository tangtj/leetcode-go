package leetcode

func hammingDistance(x int, y int) int {
	c := x ^ y
	r := 0
	for ; c > 0; c = c & (c - 1) {
		r++
	}
	return r
}
