# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
      def recursive(r: 'TreeNode') -> 'TreeNode':
        if not r:
          return r
        
        if p.val < r.val and q.val < r.val:
          return recursive(r.left)
        
        if p.val > r.val and q.val > r.val:
          return recursive(r.right)
        
        return r
          
      return recursive(root)
        