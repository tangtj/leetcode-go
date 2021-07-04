package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	l, r := 0, n
	target := 0
	for l <= r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			target = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return target
}
