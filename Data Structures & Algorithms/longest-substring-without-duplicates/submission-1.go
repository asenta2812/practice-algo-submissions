func lengthOfLongestSubstring(s string) int {
	window := make(map[rune]bool)
	l := 0
	res := 0

	for r := 0; r < len(s); r++ {
		for window[rune(s[r])] {
			delete(window, rune(s[l]))
			l++
		}
		window[rune(s[r])] = true 
		res = max(r - l + 1, res)
	}

	return res
}