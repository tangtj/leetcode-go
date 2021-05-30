package leetcode

// n & (n-1) 可以 去掉 二进制最后一个1
// 2 的n次幂的 二进制上只有 1个 1
func isPowerOfTwo(n int) bool {

	return n > 0 && (n&(n-1) == 0)

}
