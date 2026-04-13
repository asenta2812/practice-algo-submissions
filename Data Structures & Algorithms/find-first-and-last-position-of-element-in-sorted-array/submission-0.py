class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
      if len(nums) == 0:
        return [-1, -1]
        
      lower = self.lower_bound(nums, target)
      upper = self.upper_bound(nums, target)
      return [lower, upper]
    

    def lower_bound(self, nums: List[int], target: int) -> int:
      left, right = 0, len(nums) - 1
      while left < right:
        mid = (left + right) // 2
        if nums[mid] > target:
          right = mid - 1
        elif nums[mid] < target:
          left = mid + 1
        else:
          right = mid
      return left if nums[left] == target else -1

    def upper_bound(self, nums: List[int], target: int) -> int:
      left, right = 0, len(nums) - 1
      while left < right:
        mid = (left + right) // 2 + 1 # +1 avoid infinite loop
        if nums[mid] > target:
          right = mid - 1
        elif nums[mid] < target:
          left = mid + 1
        else:
          left = mid 
      return right if nums[right] == target else -1
        