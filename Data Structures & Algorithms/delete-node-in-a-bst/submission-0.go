/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Right)
			root.Val = minNode.Val
			// remove minNode
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}	

	return root
}

func minValueNode(root *TreeNode) *TreeNode {
	cur := root
	for root != nil && cur.Right != nil {
		cur = cur.Right
	}

	return cur
}
