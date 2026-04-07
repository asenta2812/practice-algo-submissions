class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
		let l = 0, 
			r = s.length - 1

		while (l < r) {
			if (!this.alphaNum(s[l])) {
				l++
				continue
			}
			if (!this.alphaNum(s[r])) {
				r--
				continue
			}

			if (s[l].toLowerCase() !== s[r].toLowerCase()) {
				return false
			}
			l++
			r--
		}

		return true
	}

	alphaNum(c) {
		return (
			(c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9')
		)
	}
}
