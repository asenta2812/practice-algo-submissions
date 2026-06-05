class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        n = len(nums)
        low, high = 1, n - 1
        while low < high:
            mid = low + (high - low) // 2
            less_or_equal = sum(1 for num in nums if num <= mid)
            if less_or_equal <= mid:
                low = mid + 1
            else:
                high = mid
        
        return low