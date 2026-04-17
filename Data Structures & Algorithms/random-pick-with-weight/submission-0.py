class Solution:

    def __init__(self, w: List[int]):
      self.prefix_sums = [w[0]]
      for i in range(1, len(w)):
        self.prefix_sums.append(self.prefix_sums[-1] + w[i])

    def pickIndex(self) -> int:
      target = random.randint(1, self.prefix_sums[-1])

      l, r = 0, len(self.prefix_sums) - 1
      while l < r:
        m = (l+r)//2
        if self.prefix_sums[m] < target:
          l = m + 1
        else:
          r = m
      return l
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()