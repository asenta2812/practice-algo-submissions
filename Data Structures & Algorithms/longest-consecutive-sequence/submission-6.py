class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
       
        if not nums:
            return 0

        res = 0
        s = set(nums)

        for num in s:
            if not num - 1 in s:
                current_num = num
                count = 1

                while current_num + 1 in s:
                    current_num += 1
                    count += 1
                res = max(res, count)
                

        return res