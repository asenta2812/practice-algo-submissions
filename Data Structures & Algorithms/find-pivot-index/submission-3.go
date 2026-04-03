func pivotIndex(nums []int) int {
	preArr := make([]int, len(nums))
	pre := 0
	for i, v := range nums {
		preArr[i] = pre + v
		pre += v
	}

	for i := 0; i < len(nums); i++ {
		r1 := 0 
		if i > 0 {
			r1 = preArr[i-1]
		}
		
		if r1 == preArr[len(nums) - 1] - preArr[i] {
			return i
		}
	}

	return -1
}