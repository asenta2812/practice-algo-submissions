class Solution {
    /**
     * @param {string[]} sentence1
     * @param {string[]} sentence2
     * @param {string[][]} similarPairs
     * @return {boolean}
     */
    areSentencesSimilar(sentence1, sentence2, similarPairs) {
        const m = new Map()
        if (sentence1.length !== sentence2.length) {
            return false
        }

        for (let [pair1, pair2] of similarPairs) {
            const key = this.getTemplate(pair1, pair2)
            m.set(key, true)
        }

        for (let i in sentence1) {
            if (sentence1[i] === sentence2[i]) {
                continue
            }

            const key = this.getTemplate(sentence1[i], sentence2[i])
            if (m.has(key)) {
                 continue
            }

            return false
        }

        return true
    }

    getTemplate(str1, str2) {
        if (str1 > str2) {
            return `${str1}:${str2}`
        }

        return `${str2}:${str1}`
    }
}
