# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        count = 0

        def dfs(node, max_prev):
            nonlocal count
            if not node:
                return

            if node.val >= max_prev:
                count+=1

            dfs(node.left, max(node.val, max_prev))
            dfs(node.right, max(node.val, max_prev))
            
        dfs(root, root.val)
        return count
        