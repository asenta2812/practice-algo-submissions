type NumMatrix struct {
	sumArray [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	sumArray := make([][]int, rows)
	for i := 0; i < rows; i++ {
		prefix := 0
		for j := 0; j < cols; j++ {
			sumArray[i] = append(sumArray[i], matrix[i][j] + prefix)
			prefix += matrix[i][j]
		}
	}

	fmt.Println(sumArray)
	return NumMatrix{
		sumArray: sumArray,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
    for row := row1; row <= row2; row++ {
        if col1 > 0 {
            res += this.sumArray[row][col2] - this.sumArray[row][col1-1]
        } else {
            res += this.sumArray[row][col2]
        }
    }
    return res
}



// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
