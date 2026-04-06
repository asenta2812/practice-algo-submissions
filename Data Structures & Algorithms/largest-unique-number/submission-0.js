class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    largestUniqueNumber(nums) {
        const m = {}

        for (let v of nums) {
            m[v] = (m[v] || 0) + 1
        }

        let max = -1
        for (let v of Object.keys(m)) {
            if (m[v] === 1) {
                max = Math.max(max, Number(v))
            }  
        }

        return max
    }
}
