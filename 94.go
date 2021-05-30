package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return deepin(root)
}

func deepin(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	num := make([]int, 0)
	if root.Left != nil {
		num = append(num, deepin(root.Left)...)
	}
	num = append(num, root.Val)
	if root.Right != nil {
		num = append(num, deepin(root.Right)...)
	}
	return num
}
