/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil || targetSum == 0 {
		return false
	}

	var dfs func(node *TreeNode, target int) bool
	dfs = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}

		// logic node.Val
		target -= node.Val
		if target == 0 && node.Left == nil && node.Right == nil {
			return true
		}
		return dfs(node.Left, target) || dfs(node.Right, target)
	}

	return dfs(root, targetSum)
}
