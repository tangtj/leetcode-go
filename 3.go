package leetcode

func lengthOfLongestSubstring(s string) int {
	cmap := make(map[rune]int, 0)
	max := 0

	start := 0
	for i, c := range []rune(s) {
		if idx, ok := cmap[c]; ok {
			start = MathMax(start, idx)
		}
		max = MathMax(max, i-start+1)
		cmap[c] = i + 1

	}
	return max
}
