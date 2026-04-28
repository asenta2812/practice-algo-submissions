class Solution:
    def maxProduct(self, nums: List[int]) -> int:
      if not nums: 
        return 0
      
      res = cur_max = cur_min = nums[0]
      for num in nums[1:]:
        opt1 = num
        opt2 = num * cur_max
        opt3 = num * cur_min

        cur_max = max(opt1, opt2, opt3)
        cur_min = min(opt1, opt2, opt3)
        res = max(res, cur_max)


      return res