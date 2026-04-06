class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
    isValidSudoku(board) {
        const rows = {}
        const cols = {}
        const squares = {}
        const n = 9

        for (let r = 0; r < n; r++) {
            for (let c = 0; c < n; c++) {
                const v = board[r][c]
                if (v === ".") {
                    continue
                }

                // 0:0, 0:1, 0:2
                const squareKey = `${Math.floor(r / 3)}:${Math.floor(c / 3)}`

                if (
                    rows[r]?.includes(v) ||
                    cols[c]?.includes(v) ||
                    squares[squareKey]?.includes(v)
                ) {
                    return false
                }

                // set
                if (!rows[r]) {
                    rows[r] = []
                }
                if (!cols[c]) {
                    cols[c] = []
                }
                if (!squares[squareKey]) {
                    squares[squareKey] = []
                }

                rows[r].push(v)
                cols[c].push(v)
                squares[squareKey].push(v)
            }
        }

        return true
    }
}
