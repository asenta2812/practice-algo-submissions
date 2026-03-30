func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	l := 0
	currSum := 0

	for r := 0; r < len(nums); r++ {
		currSum += nums[r]
		for currSum >= target {
			diff := r - l + 1
			res = min(diff, res)
			currSum -= nums[l]
			l++
		}
	}

	if res == len(nums) + 1 {
		return 0
	}

	return res
}
