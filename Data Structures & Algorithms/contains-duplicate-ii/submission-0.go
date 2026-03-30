func containsNearbyDuplicate(nums []int, k int) bool {
	l := 0
	window := make(map[int]bool)
	for r := 0; r < len(nums); r++ {
		diff := r - l 
		if diff > k {
			delete(window, nums[l])
			l++
		}

		if _, ok := window[nums[r]]; ok {
			return true
		}

		window[nums[r]] = true
	}

	return false
}
