type NumArray struct {
    sumArray []int
}


func Constructor(nums []int) NumArray {
    prefix := 0
	sumArray := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumArray[i] = prefix + nums[i]
		prefix += nums[i]
	}

	return NumArray {
		sumArray: sumArray,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
		return this.sumArray[right]
	}

	return this.sumArray[right] - this.sumArray[left-1]
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

