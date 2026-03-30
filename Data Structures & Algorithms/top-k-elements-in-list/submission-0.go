
func topKFrequent(nums []int, k int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    // map item with frequency
    mItemFreq := make(map[int]int)
    for _, v := range nums {
        mItemFreq[v]++
    }
    // array key: frequency, value is array item
    arrFreIems := make([][]int, len(nums) + 1)

    for item, freq := range mItemFreq {
        arrFreIems[freq] = append(arrFreIems[freq], item)
    }

    // using bucket sort
    result := []int{}
    for i := len(arrFreIems) - 1; i > 0; i-- {
        if len(arrFreIems[i]) == 0 {
            continue
        }
        for _, v := range arrFreIems[i] {
            result = append(result, v)
            if len(result) == k {
                return result
            }
        }
    }

    return []int{}
}
