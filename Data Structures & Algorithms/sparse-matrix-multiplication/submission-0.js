class Solution {
    /**
     * @param {number[][]} mat1
     * @param {number[][]} mat2
     * @return {number[][]}
     */
    multiply(mat1, mat2) {
        const m = mat1.length
        const k = mat1[0].length
        const n = mat2[0].length

        let result = Array.from({length: m}, () => new Array(n).fill(0))


      // p =  [0, k - 1]  => sum( mat1[i][p] * mat2[p][j]) 
      for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let p = 0; p < k; p++) {
                result[i][j] += mat1[i][p] * mat2[p][j]
            }
        }
      } 

      return result
    }
}
