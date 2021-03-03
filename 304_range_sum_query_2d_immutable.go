package leetcode

type NumMatrix struct {
	sums [][]int
}

func Constructor2(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		sum := make([]int, len(matrix[i]))

		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 && j == 0 {
				sum[j] = matrix[i][j]
				continue
			}

			sum[j] = matrix[i][j]
			if j > 0 {
				sum[j] += sum[j-1]
			}
			if i > 0 {
				sum[j] += sums[i-1][j]
			}
			if i > 0 && j > 0 {
				sum[j] -= sums[i-1][j-1]
			}
		}

		sums[i] = sum
	}

	return NumMatrix{
		sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.sums[row2][col2]
	if col1 > 0 {
		sum = sum - this.sums[row2][col1-1]
	}
	if row1 > 0 {
		sum = sum - this.sums[row1-1][col2]
	}
	if col1 > 0 && row1 > 0 {
		sum += this.sums[row1-1][col1-1]
	}
	return sum
}
