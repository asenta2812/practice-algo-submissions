func isPalindrome(s string) bool {
	s = CleanString(s)
	L, R := 0, len(s) - 1
	for L < R {
		if s[L] != s[R] {
			fmt.Println(string(s[L]),string(s[R]))
			return false
		}
		L++
		R--
	}

	return true
}
func CleanString(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return str
	}

	// 2. Remove the unwanted characters
	processed := reg.ReplaceAllString(str, "")
	
	// 3. Convert the remaining string to lowercase
	return strings.ToLower(processed)
}