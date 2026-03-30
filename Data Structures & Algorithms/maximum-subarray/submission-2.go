func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currSum, maxSum := nums[0], nums[0]

	for r := 1; r < len(nums); r++{

		currSum = max(currSum + nums[r], nums[r])
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
