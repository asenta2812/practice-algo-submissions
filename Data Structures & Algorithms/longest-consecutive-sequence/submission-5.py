class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
       
        """
        *Imagine each number as a brick. When you place the num brick down:
            It can connect a sequence on the left (ending at num - 1).
            It can connect a sequence on the right (starting at num + 1).
        """
        if not nums:
            return 0

        res, m = 0, defaultdict(int)

        for num in nums:
            # check duplicate
            if m[num]:
                continue
            left = m[num-1]
            right = m[num+1]
            new_len = left + right + 1
            res = max(new_len, res)
            m[num] = new_len
            # set 2 heads
            m[num-left] = new_len
            m[num+right] = new_len

        return res