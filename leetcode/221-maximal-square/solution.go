package maximalsquare

func MaximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])

	var max int
	var prev, curr []int

	curr = make([]int, rows)
	for i := rows - 1; i >= 0; i-- {
		prev, curr = curr, make([]int, cols)
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				continue
			}

			switch {
			case i == rows-1, j == cols-1:
				curr[j] = 1
			default:
				curr[j] = 1 + min(curr[j+1], prev[j], prev[j+1])
			}

			if max < curr[j] {
				max = curr[j]
			}
		}
	}

	return max * max
}

func min(right, down, diagonal int) int {
	min := right
	if min > down {
		min = down
	}
	if min > diagonal {
		min = diagonal
	}
	return min
}
