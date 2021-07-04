package leetcode

var m = make(map[int]int, 0)

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if r, ok := m[n]; ok {
		return r
	}
	ret := climbStairs(n-1) + climbStairs(n-2)
	m[n] = ret
	return ret
}

func climbStairs01(n int) int {
	if n <= 2 {
		return n
	}
	q, p, r := 1, 2, 3
	for i := 3; i < n; i++ {
		q, p = p, r
		r = q + p
	}
	return r
}
