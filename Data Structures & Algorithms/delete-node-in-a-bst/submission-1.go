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
			maxNode := maxValueNode(root.Right)
			root.Val = maxNode.Val
			// remove maxNode
			root.Right = deleteNode(root.Right, maxNode.Val)
		}
	}	

	return root
}

func maxValueNode(root *TreeNode) *TreeNode {
	cur := root
	for root != nil && cur.Left != nil {
		cur = cur.Left
	}

	return cur
}
