class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        

        # dp[i]: is the length of the longest ascending 
        # subsequence ending at index (ith)
        # n = len(nums)
        # dp = [1] * (n)
        # res = 1

        # for i in range(n):
        #     for j in range(i):
        #         if nums[i] > nums[j]:
        #             dp[i] = max(dp[i], dp[j] + 1)
        #     res = max(res, dp[i])
        # return res

        # binary search
        def binary_search(arr: List[int], target: int) -> int:
            if not arr:
                return None
            
            l, r = 0, len(arr)
            while l < r:
                m = (l+r)//2
                if arr[m] < target:
                    l = m + 1
                else:
                    r = m
            return l if l < len(arr) else None

        arr = []

        for num in nums:
            idx = binary_search(arr, num)
            if idx != None:
                arr[idx] = num
            else:
                arr.append(num)
        return len(arr)




