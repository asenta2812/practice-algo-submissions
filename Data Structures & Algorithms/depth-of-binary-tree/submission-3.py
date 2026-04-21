# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
  def maxDepth(self, root: Optional[TreeNode]) -> int:
    if not root:
      return 0


    # res = 1
    #dfs
    # stack = [[root, 1]]

    # while stack:
    #   node, depth = stack.pop()
    #   # logic here
    #   res = max(depth, res)

    #   if node.left:
    #     stack.append([node.left, depth + 1])
    #   if node.right:
    #     stack.append([node.right, depth + 1])

    # bfs
    q = deque()
    q.append(root)
    level = 0

    while q:
      for _ in range(len(q)):
        node = q.popleft()
        if node.left:
          q.append(node.left)
        if node.right:
          q.append(node.right)
      level += 1

    return level
