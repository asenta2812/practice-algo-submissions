# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:   
    def serialize(self, node: Optional[TreeNode]):
      if not node:
        return "#"
      left = self.serialize(node.left)
      right = self.serialize(node.right)
      return f":{node.val}:{left},{right}"
      
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
      root_sl = self.serialize(root)
      sub_sl = self.serialize(subRoot)
      return sub_sl in root_sl