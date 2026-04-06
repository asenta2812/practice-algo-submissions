class Solution {
    /**
     * @param {number[][]} arrays
     * @return {number}
     */
    maxDistance(arrays) {
        let res = 0
        let minVal = arrays[0][0]
        let maxVal = arrays[0][arrays[0].length - 1]

        for (let i = 1; i < arrays.length; i++) {
            const array = arrays[i]
            const len = array.length

            const d1 = Math.abs(array[len - 1] - minVal)
            const d2 = Math.abs(maxVal - array[0])
            res = Math.max(
                res,
                Math.max(d1, d2)
            )
            minVal = Math.min(minVal, array[0])
            maxVal = Math.max(maxVal, array[len - 1])
        }

        return res
    }
}

123
456
789

13
16

