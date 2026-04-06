class Solution {
    /**
     * @param {number[]} arr
     * @return {number[]}
     */
    replaceElements(arr) {
        const n = arr.length
        let rightMax = -1
        for (let i = n - 1; i >= 0; i--) {
            const crr = arr[i]
            arr[i] = rightMax
            rightMax = Math.max(crr, rightMax)
        }

        return arr
    }
}
