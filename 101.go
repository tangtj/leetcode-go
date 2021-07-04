package leetcode

func isSymmetric(root *TreeNode) bool {
	return dfs101(root.Left, root.Right)
}

func dfs101(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return dfs101(l.Left, r.Right) && dfs101(l.Right, r.Left)
}
