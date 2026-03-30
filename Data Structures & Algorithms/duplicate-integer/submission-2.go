
import (
    "slices"
)

func hasDuplicate(nums []int) bool {
    slices.Sort(nums) 

    for i, v := range nums[1:] {
        if v == nums[i] {
            return true
        }
    }

    return false
}
