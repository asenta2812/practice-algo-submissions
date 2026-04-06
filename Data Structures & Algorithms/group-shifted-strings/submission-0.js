class Solution {
    /**
     * @param {string[]} strings
     * @return {string[][]}
     */
    groupStrings(strings) {
        function getHash(str) {
            const key = []

            for (let i = 0; i < str.length - 1; i++) {
                const diff = str.charCodeAt(i + 1) - str.charCodeAt(i)
                const round = (diff + 26) % 26
                key.push(String.fromCharCode(round + 'a'.charCodeAt(0)))
            }
            
            const res = key.join('')
            return res.length > 0 ? res : '(empty)'
        }

        const m = {}
        
        for (let str of strings) {
            const key = getHash(str)
            if (!m[key]) {
                m[key] = []
            }
            m[key].push(str)
        }

        console.log(m)

        return Object.values(m)
    }
}
