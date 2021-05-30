package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var ret = &ListNode{}
	var prev = ret
	if l1 == nil && l2 == nil {
		return ret.Next
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return ret.Next

}
