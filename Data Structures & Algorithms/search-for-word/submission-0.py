class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        
        # 1.choise

        # 2.goal

        # 3.contraint
        rows, cols = len(board), len(board[0])
        path = set()
        def dfs(r: int, c: int, i: int) -> bool:
            if i == len(word):
                return True
                
            if r >= rows or c >= cols or r < 0 or c < 0 or board[r][c] != word[i] or (r,c) in path:
                return False
            
            path.add((r,c))
            res = (
                dfs(r+1, c, i+1) or
                dfs(r-1,c, i+1) or
                dfs(r, c+1,i+1) or
                dfs(r, c-1, i+1)
            )
            if res:
                print((r,c))
            path.remove((r,c))
            return res
        
        for i in range(rows):
            for j in range(cols):
                if dfs(i, j, 0):
                    return True
        return False