/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

    res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		ql := len(q)

		for i := 0; i < ql; i++ {
			c := q[0]
			q = q[1:]
			if i == ql - 1 {
				res = append(res, c.Val)
			}

			if c.Left != nil {
				q = append(q, c.Left)
			}
			if c.Right != nil {
				q = append(q, c.Right)
			}
		}
	}
	return res
}
