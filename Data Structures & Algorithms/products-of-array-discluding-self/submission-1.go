func productExceptSelf(nums []int) []int {
    if len(nums) == 0 {
        return nums
    }
    res := make([]int, len(nums))

    // a, b := 1, 1

    // for i, v := range nums {
    //     if v != 0 {
    //         b *= v
    //     }

    //     if v == 0 && a == 0 && i != 0 {
    //         b *= v
    //     }
    //     a *= v
    // }

    // for i, v := range nums {
    //     if v == 0 {
    //         res[i] = b
    //     } else {
    //         res[i] = a / v
    //     }
    // }

    // return res

   // 1 2 4 6
   // for1 = 1 1 2 8
   // for2 = 48 24 12 8
    for i,_ := range res {
        res[i] = 1
    }

    prefix := 1
    for i, v := range nums {
        res[i] = prefix
        prefix *= v
    }

    suffix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= suffix
        suffix *= nums[i]
    }
    
    return res
}

