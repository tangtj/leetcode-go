package leetcode

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return deepinR(root)
}

func deepinR(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	num := make([]int, 0)
	if root.Left != nil {
		num = append(num, deepinR(root.Left)...)
	}
	if root.Right != nil {
		num = append(num, deepinR(root.Right)...)
	}
	num = append(num, root.Val)
	return num
}
