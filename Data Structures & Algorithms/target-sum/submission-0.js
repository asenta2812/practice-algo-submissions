class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number}
     */
    findTargetSumWays(nums, target) {
        // backtrack
        const backtrack = (i, total) => {
            if(i === nums.length) {
                // 1 is correct way
                return total === target ? 1 : 0
            } 
            return (
                // plus branch
                backtrack(i + 1, total + nums[i]) + 
                // sub branch
                backtrack(i + 1, total - nums[i])
            )
        }

        return backtrack(0, 0)
    }
}
