
func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    res := 0
    for num := range numSet {
        // find the root
        if _, found := numSet[num - 1]; !found {
            currentNum, count := num, 1

            for {
                if _, ok := numSet[currentNum + 1]; ok {
                    currentNum++
                    count ++
                } else {
                    break
                }
            }
            res = max(res, count) 

        }
    }
    return res
}
