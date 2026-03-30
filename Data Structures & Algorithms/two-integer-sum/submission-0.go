func twoSum(nums []int, target int) []int {
    // for i := 0; i < len(nums); i++ {
    //     for j := i + 1; j < len(nums); j++ {
    //         if nums[i] + nums[j] == target {
    //             return []int{i, j}
    //         }
    //     }
    // }

    // return []int{}

    m := make(map[int]int)
    for i, v := range nums {
        diff := target - v
        if j, found := m[diff]; found {
            return []int{j, i}
        }
        m[v] = i
    }

    return []int{}
}
