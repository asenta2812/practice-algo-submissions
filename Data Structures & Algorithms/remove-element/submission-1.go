func removeElement(nums []int, val int) int {
	k := 0
	for _, currentValue := range nums {
		if currentValue != val {
			nums[k] = currentValue
			k++
		}
	}
	return k
}
