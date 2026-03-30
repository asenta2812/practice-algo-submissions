func removeDuplicates(nums []int) int {
    k := 1
    for index, value := range nums[1:] {
        if value != nums[index] {
            nums[k] = value
            k++
        }
    }

    return k
}
