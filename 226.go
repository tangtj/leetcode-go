package leetcode

func invertTree(root *TreeNode) *TreeNode {
	dfs226(root)
	return root
}

func dfs226(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		dfs226(root.Left)
	}
	if root.Right != nil {
		dfs226(root.Right)
	}
}
