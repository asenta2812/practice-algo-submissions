class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows_set = [set() for _ in range(9)]
        cols_set = [set() for _ in range(9)]
        squares = defaultdict(set)

        for r in range(9):
            for c in range(9):
                curr = board[r][c]
                if curr == '.':
                    continue

                if (curr in rows_set[r] or 
                    curr in cols_set[c] or
                    curr in squares[(r // 3, c // 3)] ):
                    return False 

                rows_set[r].add(curr)
                cols_set[c].add(curr)
                squares[(r // 3, c // 3)].add(curr)
        
        return True
