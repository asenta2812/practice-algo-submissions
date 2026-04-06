class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number}
     */
    findTargetSumWays(nums, target) {
       const n = nums.length 
       const dp = Array.from({length: n + 1}, () => ({}))
       dp[0][0] = 1

       for (let i = 0; i < n; i++) {
            const current = nums[i]
            for (let total in dp[i]) {
                total = Number(total)

                // plus branch
                dp[i+1][total + current] = (dp[i+1][total + current] || 0) + dp[i][total]
                // sub branch
                dp[i+1][total - current] = (dp[i+1][total - current] || 0) + dp[i][total]
            }
       }

       return dp[n][target] || 0
    }
}
