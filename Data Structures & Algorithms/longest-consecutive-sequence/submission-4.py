class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        """
            Find the first of  Consecutive Sequence (the item do not have num - 1 in this array)
        """
        if not nums:
            return 0

        res = 0

        for num in nums:
            if not num - 1 in nums:
                current_num = num
                count = 1

                while current_num + 1 in nums:
                    current_num += 1
                    count += 1
                res = max(res, count)
                

        return res