class Solution:
    def rob(self, nums: List[int]) -> int:
        # donot choose the current house: dp(i + 1)
        # choose the current house: curr_vl + dp(i + 2)
        # dp(i) = max(dp(i + 1), curr_vl + dp(i + 2))

        dp = [-1] * len(nums)

        def dps(i: int) -> int:
          if i >= len(nums):
            return 0

          if dp[i] > -1:
            return dp[i]

          dp[i] = max(nums[i] + dps(i + 2), dps(i + 1))
          return dp[i]

        return dps(0)