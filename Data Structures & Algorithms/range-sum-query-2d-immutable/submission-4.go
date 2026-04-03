type NumMatrix struct {
	sumArray [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	sumArray := make([][]int, rows)
	for row := 0; row < rows; row++ {
		prefix := 0
		for col := 0; col < cols; col++ {
            above := 0
            if row > 0 {
                above = sumArray[row-1][col]
            }
            currentSum := matrix[row][col] + prefix + above
			sumArray[row] = append(sumArray[row], currentSum)
			prefix += matrix[row][col]
		}
	}

	return NumMatrix{
		sumArray: sumArray,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    current := this.sumArray[row2][col2]
    left := 0
    if col1 > 0 {
        left = this.sumArray[row2][col1-1]
    }
    above := 0
    if row1 > 0 {
        above = this.sumArray[row1-1][col2]
    }

    insec := 0
    if col1 > 0 && row1 > 0 {
        insec = this.sumArray[row1-1][col1-1]
    }
    return current - left - above + insec
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
