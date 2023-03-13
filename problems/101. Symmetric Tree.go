package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil && r != nil || l != nil && r == nil || l.Val != r.Val {
			return false
		}
		return check(l.Right, r.Left) && check(l.Left, r.Right)
	}

	return check(root.Left, root.Right)
}
