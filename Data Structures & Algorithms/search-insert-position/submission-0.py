class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
      left, right = 0, len(nums) # because insert position can be in the end of nums

      while left < right:
        mid = (left + right) // 2
        if nums[mid] >= target:
          right = mid
        else:
          left = mid + 1
      
      return left