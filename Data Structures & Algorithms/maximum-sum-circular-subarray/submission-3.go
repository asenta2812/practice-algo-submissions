func maxSubarraySumCircular(nums []int) int {
	gloMax, currMax := nums[0], 0
	gloMin, currMin := nums[0], 0
	total := 0
	// without circular : max by kadane
	// have circular: find min, total - min (2 suffix , prefix (without min))
	for _, num := range nums {
		total += num
		
		currMax = max(currMax + num, num)
		currMin = min(currMin + num, num)
		gloMax = max(currMax, gloMax) 
		gloMin = min(currMin, gloMin)
	}
	if gloMax > 0 {
		return max(gloMax, total - gloMin)
	}

	return gloMax
}

// brute force
// res := nums[0]
// n := len(nums)

// for i := 0; i < n; i++ {
// 	currSum := 0
// 	for j := i; j < i + n; j++ {
// 		currSum += nums[j%n]
// 		res = max(currSum, res)
// 	}
// }

// return res