func isPalindrome(s string) bool {
	L, R := 0, len(s) - 1
	for L < R {
		lc := rune(s[L])
		rc := rune(s[R])

		if !(unicode.IsLetter(lc)|| unicode.IsDigit(lc)) {
			L++
			continue
		}

		if !(unicode.IsLetter(rc) || unicode.IsDigit(rc)) {
			R--
			continue
		}

		if unicode.ToLower(lc) != unicode.ToLower(rc) {
			return false
		}
		L++
		R--
	}

	return true
}