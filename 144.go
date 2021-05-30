package leetcode

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return deepinM(root)
}

func deepinM(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	num := make([]int, 0)
	num = append(num, root.Val)
	if root.Left != nil {
		num = append(num, deepinM(root.Left)...)
	}
	if root.Right != nil {
		num = append(num, deepinM(root.Right)...)
	}
	return num
}
