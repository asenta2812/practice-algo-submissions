func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currSum, maxSum := nums[0], nums[0]

	for r := 1; r < len(nums); r++{
		if currSum < 0 {
			currSum = 0
		}

		currSum += nums[r]
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
