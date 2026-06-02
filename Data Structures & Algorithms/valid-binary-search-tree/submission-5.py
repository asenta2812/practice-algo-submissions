# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True

        # BFS
        # queue = [(root, -math.inf, math.inf)]

        # while queue:
        #     curr, low, high = queue[0]
        #     queue = queue[1:]

        #     if curr.val <= low or curr.val >= high:
        #         return False

        #     if curr.left:
        #         queue.append((curr.left, low, curr.val))
        #     if curr.right:
        #         queue.append((curr.right, curr.val, high))

            
        # return True

        # DFS normal
        # def dfs(node: Optional[TreeNode], low: int, high: int) -> bool:
        #     if not node:
        #         return True
            
        #     if not (low < node.val < high):
        #         return False
            
        #     return dfs(node.left, low, node.val) and dfs(node.right, node.val, high)

        # return dfs(root, -math.inf, math.inf)

        # DFS in-order
        prev = -math.inf
        def dfs_in_order(node: Optional[TreeNode]) -> bool:
            nonlocal prev
            if not node: 
                return True
            
            if not dfs_in_order(node.left):
                return False

            # logic
            if node.val <= prev:
                return False
            
            prev = node.val

            return dfs_in_order(node.right)

        return dfs_in_order(root)












        