package leetcode

func countBits(n int) []int {
	// n&(n-1) 可以移除掉二进制上最后一个0
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t := i
		count := 0
		for ; t > 0; t = t & (t - 1) {
			count++
		}
		r[i] = count
	}
	return r
}
