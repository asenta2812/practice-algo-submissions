class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        rows = len(matrix)
        cols = len(matrix[0])

        rows_set = set()
        cols_set = set()   

        for row in range(rows):
            for col in range(cols):
                if matrix[row][col] == 0:
                    rows_set.add(row)
                    cols_set.add(col)

        for row in range(rows):
            for col in range(cols):
                if row in rows_set or col in cols_set:
                    matrix[row][col] = 0
                    