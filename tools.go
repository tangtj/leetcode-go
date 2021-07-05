package leetcode

func ToListNode(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := &ListNode{}
	f := head
	for _, num := range nums {
		f.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		f = f.Next
	}
	return head.Next
}
