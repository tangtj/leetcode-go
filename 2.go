package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	ret := &ListNode{}
	last := ret
	for l1 != nil || l2 != nil {
		ret.Next = &ListNode{}
		ret = ret.Next
		s := 0
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		s += carry
		ret.Val, carry = s%10, s/10
	}
	if carry > 0 {
		ret.Next = &ListNode{Val: carry}
	}
	return last.Next
}
