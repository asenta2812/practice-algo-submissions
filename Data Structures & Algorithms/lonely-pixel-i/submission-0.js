class Solution {
    /**
     * @param {character[][]} picture
     * @return {number}
     */
    findLonelyPixel(picture) {
        let count = 0
        const colLen = picture.length
        const rowLen = picture[0].length
        let countColumn = new Array(colLen).fill(0)
        let countRow = new Array(rowLen).fill(0)
        for (let col = 0; col < colLen; col++) {
            for (let row = 0; row < rowLen; row++) {
                if (picture[col][row] === "B") {
                    countColumn[col]++
                    countRow[row]++
                }
            }
        }

        for (let col = 0; col < colLen; col++) {
            for (let row = 0; row < rowLen; row++) {
                if (picture[col][row] === "B" && countColumn[col] == 1 && countRow[row] == 1) {
                   count++
                }
            }
        } 

        return count
    }
}
