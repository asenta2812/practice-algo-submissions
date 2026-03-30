func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	queue := [][]int{{0,0}}
	grid[0][0] = 1

	directions := [][]int{
			{-1,-1}, {-1,0}, {-1,1}, 
	         {0, -1},        {0,1}, 
			{1, -1}	,{1, 0}, {1,1}}

	length := 1
	for len(queue) > 0 {
		lenQueueAtLevel := len(queue)
		for i:=0 ; i < lenQueueAtLevel; i++ {
			cur := queue[0]
			r, c := cur[0], cur[1]
			queue = queue[1:]
			
			for _, v := range directions {
				dr, dc := r + v[0], c + v[1]
				if dr >= 0 && dc >= 0 && dr < n && dc < n && grid[dr][dc] == 0 {
                    if dr == n - 1 && dc == n - 1 {
                        return length + 1
                    }
                    grid[dr][dc] = 1 // Đánh dấu hàng xóm đã thăm
                    queue = append(queue, []int{dr, dc})
                }
			}
		}
		length++
	}

	return length
}

