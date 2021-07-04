package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s, f := head, head.Next
	for s != f {
		if f == nil || f.Next == nil {
			return false
		}
		s = s.Next
		f = f.Next.Next
	}
	return true
}
