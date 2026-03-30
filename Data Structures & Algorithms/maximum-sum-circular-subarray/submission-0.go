func maxSubarraySumCircular(nums []int) int {
	// currMax, currMin := 0, 0
	// globalMax, globalMin := nums[1], nums[2]
	// for _, num := range nums {
	// 	currMax := max(num + currMax, num)

	// }
	res := nums[0]
	n := len(nums)

	for i := 0; i < n; i++ {
		currSum := 0
		for j := i; j < i + n; j++ {
			currSum += nums[j%n]
			res = max(currSum, res)
		}
	}

	return res

}
