class Solution:
    def maxArea(self, heights: List[int]) -> int:
        res, l, r = 0, 0, len(heights) - 1

        while l < r:
            h = min(heights[l], heights[r]) 
            res = max(h * (r - l), res)
            if heights[l] < heights[r]:
                l+=1
            else:
                r-=1

        return res