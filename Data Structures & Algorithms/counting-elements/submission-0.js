class Solution {
    /**
     * @param {number[]} arr
     * @return {number}
     */
    countElements(arr) {
        const m = {}
        for (let v of arr) {
            m[v] = (m[v] || 0) + 1
        }

        console.log(m)

        let count = 0
        for (let v of Object.keys(m)) {
            if (m[Number(v) + 1]) {
                count += m[v]
            }
        }

        return count
    }
}
