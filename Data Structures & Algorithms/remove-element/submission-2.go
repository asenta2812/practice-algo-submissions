// func removeElement(nums []int, val int) int {
// 	k := 0
// 	for _, currentValue := range nums {
// 		if currentValue != val {
// 			nums[k] = currentValue
// 			k++
// 		}
// 	}
// 	return k
// }

func removeElement(nums []int, val int) int {
	i, l := 0, len(nums)
	for i < l {
		if nums[i] == val {
			l--
			nums[i] = nums[l]
		} else {
			i++
		}
	}
	return i
}