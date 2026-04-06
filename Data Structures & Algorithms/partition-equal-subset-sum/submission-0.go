func canPartition(nums []int) bool {
	total := 0 
	for _, v := range nums {
		total += v
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2
	dp := make([]bool, target + 1)
	dp[0] = true
	for _, num := range nums {
		for c := target; c >= num; c-- {
			dp[c] = dp[c] || dp[c-num]
		}
	}

	return dp[target]
}

