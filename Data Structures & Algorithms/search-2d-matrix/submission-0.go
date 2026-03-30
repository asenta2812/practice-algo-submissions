// func searchMatrix(matrix [][]int, target int) bool {
// 	row, col := len(matrix), len(matrix[0])
// 	le := row*col

// 	flatlist := make([]int, 0 , le)
// 	for _, list := range matrix {
// 		flatlist = append(flatlist, list...)
// 	}

// 	l, r := 0, le - 1
// 	for l <= r {
// 		m := (l+r)/2
// 		if target < flatlist[m] {
// 			r = m - 1
// 		} else if target > flatlist[m] {
// 			l = m + 1
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])
	low, high := 0, rows*cols - 1
	for low <= high {
		mid := (low+high)/2
		r := mid / 2
		c := mid % 2
		midValue := matrix[r][c]
		if midValue == target {
			return true
		} else if target < midValue {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
