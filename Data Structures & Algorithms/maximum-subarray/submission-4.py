class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
      curr_sum, max_sum = nums[0], nums[0]

      for v in nums[1:]:
        curr_sum = max(curr_sum + v, v)
        max_sum = max(curr_sum, max_sum)
      return max_sum