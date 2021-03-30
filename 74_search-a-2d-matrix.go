package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	// determine row first
	// then determine column
	lastRow := len(matrix) - 1
	for i := 0; i < lastRow; i++ {
		// row
		if matrix[i][0] <= target && target < matrix[i+1][0] {
			// search in this row
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				}
			}
			return false
		}
	}
	// try to search in the last row
	for j := 0; j < len(matrix[lastRow]); j++ {
		if matrix[lastRow][j] == target {
			return true
		}
	}
	return false
}
