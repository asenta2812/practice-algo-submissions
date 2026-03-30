/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func insertIntoBST(root *TreeNode, val int) *TreeNode {
//     if root == nil {
// 		return &TreeNode{
// 			Val: val,
// 		}
// 	}

// 	if val > root.Val {
// 		root.Right = insertIntoBST(root.Right, val)
// 	} else if val < root.Val {
// 		root.Left = insertIntoBST(root.Left, val)
// 	}
// 	return root
// }

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}

	if root == nil {
		return newNode
	}

	cur := root
	for {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = newNode
				return root
			}
			cur = cur.Right
		} else {
			if cur.Left == nil {
				cur.Left = newNode
				return root
			}
			cur = cur.Left
		}
	}


}