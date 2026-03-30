/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	stack := []*TreeNode{}
	cur := root 
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// top of stack
		cur = stack[len(stack) - 1]
		// pop stack
		stack = stack[:len(stack)-1]
		arr = append(arr, cur)
		cur = cur.Right
	}

	return arr
}
