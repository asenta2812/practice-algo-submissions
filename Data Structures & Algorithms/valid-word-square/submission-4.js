class Solution {
    /**
     * @param {string[]} words
     * @return {boolean}
     */
    validWordSquare(words) {
        for (let columnWord = 0; columnWord < words.length; columnWord++) {
            for (let rowWord = 0; rowWord < words[columnWord].length; rowWord++) {
                if (
                    rowWord >= words.length || 
                    words[columnWord][rowWord] !== words[rowWord][columnWord]) {
                    return false
                }
            }
        }

        return true
    }
}

// abc
// b