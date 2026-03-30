func numOfSubarrays(arr []int, k int, threshold int) int {
	res, l := 0, 0
	currSum := 0
	for r := 0; r < len(arr); r++ {
		diff := r - l + 1
		currSum += arr[r]

		if diff < k {
			continue
		}

		average := currSum / k
		if average >= threshold {
			res++
		}
		currSum -= arr[l]
		l++
	}

	return res
}
