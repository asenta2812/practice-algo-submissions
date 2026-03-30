func numIslands(grid [][]byte) int {
    if len(grid) == 0 { return 0 }

	count := 0
	for ir, row := range grid {
		for ic, _ := range row {
			if grid[ir][ic] == '1' {
				count ++ 
				dfs(grid, ir, ic)
			}
		}
	}

	return count
}

func dfs(grid[][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	// logic in current land
	// remove land to water
	grid[r][c] = '0'

	dfs(grid, r - 1, c) // up
	dfs(grid, r + 1, c) // down
	dfs(grid, r, c - 1) // left
	dfs(grid, r, c + 1)

} 
