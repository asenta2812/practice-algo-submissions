class Solution {
    /**
     * @param {number[]} nums1
     * @param {number[]} nums2
     * @return {number[]}
     */
    anagramMappings(nums1, nums2) {
        const m = new Map()
        for (let i in nums2) {
            m.set(nums2[i], i)
        }

        const ans = new Array(nums1.length)
        for (let i in nums1) {
            ans[i] = m.get(nums1[i])
        }

        return ans
    }
}
