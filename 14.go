package leetcode

func longestCommonPrefix(strs []string) string {
	if i := len(strs); i <= 0 {
		return ""
	} else if i == 1 {
		return strs[0]
	}

	var target uint8 = 0
	ret := make([]uint8, 0)
	idx := 0

	for true {
		for i, j := 0, len(strs); i < j; i++ {
			t := strs[i]
			if idx >= len(t) {
				return string(ret)
			} else if i == 0 {
				target = t[0]
				strs[i] = t[1:]
				continue
			} else if target != t[0] {
				return string(ret)
			} else if i == j-1 {
				ret = append(ret, t[0])
				strs[i] = t[1:]
			} else {
				strs[i] = t[1:]
				continue
			}
		}
	}
	return string(ret)
}
