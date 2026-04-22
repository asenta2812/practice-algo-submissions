class Solution:
    def rob(self, nums: List[int]) -> int:
        # donot choose the current house: dp(i + 1)
        # choose the current house: curr_vl + dp(i + 2)
        # dp(i) = max(dp(i + 1), curr_vl + dp(i + 2))
        # top-down

        # dp = [-1] * len(nums)

        # def dps(i: int) -> int:
        #   if i >= len(nums):
        #     return 0

        #   if dp[i] > -1:
        #     return dp[i]

        #   dp[i] = max(nums[i] + dps(i + 2), dps(i + 1))
        #   return dp[i]

        # return dps(0)

        # bottom up
        # max max(nums[i] + dps(i - 2), dps(i - 1))
        # if not nums:
        #   return 0
        # if len(nums) == 1:
        #   return nums[0]

        # dp = [-1] * len(nums)
        # dp[0] = nums[0]
        # dp[1] = max(nums[1], dp[0])
        # for i in range(2, len(nums)):
        #   dp[i] = max(nums[i] + dp[i - 2], dp[i-1])
        # return dp[len(nums) - 1]
        rob1, rob2 = 0, 0
        for num in nums:
          tmp = max(num + rob2, rob1)
          rob2 = rob1
          rob1 = tmp
        return rob1


