class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        l, r = 0, m * n - 1
        while l <= r:
          mid = (l + r) // 2
          row, col = mid // n, mid % n
          curr = matrix[row][col]
          if curr == target:
            return True
          elif curr > target:
            r = mid - 1
          else:
            l = mid + 1
        return False