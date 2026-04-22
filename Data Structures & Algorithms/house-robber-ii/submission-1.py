class Solution:
    def rob(self, nums: List[int]) -> int:
      # 1 -> n - 1: dont rob first house => take last house
      # 0 -> n - 2: dont rob last house => takse first house 
      if not nums:
        return 0
      if len(nums) == 1:  return nums[0]

      def solve(sub_nums: List[int]) -> int:
        rob1, rob2 = 0, 0
        for num in sub_nums:
          tmp = max(num + rob2, rob1)
          rob2 = rob1 
          rob1 = tmp 
        return rob1

      return max(solve(nums[1:]), solve(nums[:-1]))