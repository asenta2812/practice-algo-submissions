class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number}
     */
    findTargetSumWays(nums, target) {
       let dp = new Map()
       dp.set(0, 1)

       for (let num of nums) {
            let newDp = new Map()
            for (let [total, count] of dp) {
                // plus branch
                // dp[i+1][total + current] = (dp[i+1][total + current] || 0) + dp[i][total]
                newDp.set(total + num, (newDp.get(total + num) || 0) + count)
                // sub branch
                newDp.set(total - num, (newDp.get(total - num) || 0) + count)
            }
            dp = newDp
       }

       return dp.get(target) || 0
    }
}
