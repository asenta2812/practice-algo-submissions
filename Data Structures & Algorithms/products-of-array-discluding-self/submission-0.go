func productExceptSelf(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }

    a, b := 1, 1
    res := make([]int, len(nums))

    for i, v := range nums {
        if v != 0 {
            b *= v
        }

        if v == 0 && a == 0 && i != 0 {
            b *= v
        }
        a *= v
    }

    for i, v := range nums {
        if v == 0 {
            res[i] = b
        } else {
            res[i] = a / v
        }
    }

    return res
}

