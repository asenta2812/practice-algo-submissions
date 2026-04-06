class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
        const n = nums.length
        if (!n) {
            return 0
        }

        const m = {}
        let res = 0
        for (let num of nums) {
            if (m[num]) {
                continue
            }

            /*
            *Hãy tưởng tượng mỗi con số là một viên gạch. Khi bạn đặt viên gạch num xuống:      
                Nó có thể nối một dãy bên trái (kết thúc tại num - 1).
                Nó có thể nối một dãy bên phải (bắt đầu tại num + 1).
            */
            const left = m[num-1] || 0
            const right = m[num+1] || 0
            const newLen = left + right + 1
            res = Math.max(newLen, res)
            m[num] = res
            m[num-left] = newLen
            m[num+right] = newLen
        }

        return res
    }
}
