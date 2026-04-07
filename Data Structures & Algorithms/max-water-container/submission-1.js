class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
        let res = 0,
            l = 0,
            r = heights.length - 1
        
        while (l < r) {
            let minHeight = Math.min(heights[l], heights[r])
            res = Math.max(res, minHeight * (r - l))
            if (heights[l] < heights[r]) {
                l++
            } else {
                r--
            }
        }

        return res
    }
}
