/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	// bfs
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		queuelen := len(queue)
		levels := []int{}

		for i := 0 ; i < queuelen; i++ {
			node := queue[0]
			queue = queue[1:]
			levels = append(levels, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			
		}
		res = append(res, levels)
	}

	return res
}
