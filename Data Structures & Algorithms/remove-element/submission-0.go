func removeElement(nums []int, val int) int {
    numsResult := make([]int, len(nums))
	k := 0

	for _, value := range nums {
		if value != val {
			numsResult[k] = value
			k++
		}
	}

	copy(nums, numsResult)
	return k
}
