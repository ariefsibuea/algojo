package hackerrank

// Description:
// Given a square matrix, calculate the absolute difference between sums of its diagonal.
// Example:
// 	1 2 3
// 	4 5 6
// 	9 8 9
// 	diff = |15 - 17| = 2

func DiagonalDifference(arr [][]int32) int32 {
	size := len(arr)

	var diag1, diag2 int32
	for i := 0; i < size; i++ {
		// diag1:
		// 	row = i
		// 	col = i
		// diag2:
		// 	row = i
		// 	col = row-n
		diag1 += arr[i][i]
		diag2 += arr[i][(size-1)-i]
	}

	if diag1-diag2 < 0 {
		return (diag1 - diag2) * -1
	}
	return diag1 - diag2
}
