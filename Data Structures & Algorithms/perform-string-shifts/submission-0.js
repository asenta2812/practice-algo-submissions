class Solution {
    /**
     * @param {string} s
     * @param {number[][]} shift
     * @return {string}
     */
    stringShift(s, shift) {
        // sum to one direction
        let leftShift = 0 
        for (let [direction, amount] of shift) {
            if (direction === 1) {
                leftShift -= amount
            } else {
                leftShift += amount
            }
        }

        const n = s.length
        leftShift = ((leftShift % n) + n) % n
        return s.substring(leftShift) + s.substring(0, leftShift)
    }
}
