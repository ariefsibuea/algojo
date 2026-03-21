package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SpiralMatrixII", RunTestSpiralMatrixII)
}

/*
 * Problem	: Spiral Matrix II
 * Topics	: Array, Matrix, Simulation
 * Level	: Medium
 * URL		: https://leetcode.com/problems/spiral-matrix-ii/
 *
 * Description:
 *		Create an n x n square matrix filled with numbers from 1 to n² in clockwise spiral order. The spiral pattern
 *		starts from the top-left corner and moves right across the first row, then down along the last column, then left
 *		across the bottom row, then up along the first column, and continues this spiral pattern inward until all
 *		positions are filled.
 *
 * Constraints:
 *		- 1 <= n <= 20
 *
 * Examples:
 *		Example 1:
 *		Input: n = 3
 *		Output: [[1,2,3],[8,9,4],[7,6,5]]
 *		Explanation: The matrix is filled in spiral order: 1→2→3→4→5→6→7→8→9, starting from top-left,
 *		going right to (0,2), then down to (2,2), then left to (2,0), then up to (1,0), and finally right to (1,1).
 *
 *		Example 2:
 *		Input: n = 1
 *		Output: [[1]]
 *		Explanation: A single cell matrix containing the value 1.
 */

func generateMatrix(n int) [][]int {
	return generateMatrixIterative(n)
}

func generateMatrixRecursive(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // bottom
		{0, -1}, // left
		{-1, 0}, // top
	}

	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
	}

	var fillCell func(d, r, c int) bool
	var num = 1

	fillCell = func(d, r, c int) bool {
		if r < 0 || r >= n || c < 0 || c >= n {
			return false
		}
		if table[r][c] > 0 {
			return false
		}

		table[r][c] = num
		num++

		nr, nc := r+directions[d][0], c+directions[d][1]
		if !fillCell(d, nr, nc) {
			d = (d + 1) % 4
			nr, nc = r+directions[d][0], c+directions[d][1]
			fillCell(d, nr, nc)
		}

		return true
	}

	fillCell(0, 0, 0)
	return table
}

func generateMatrixIterative(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // bottom
		{0, -1}, // left
		{-1, 0}, // top
	}

	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
	}

	d, r, c := 0, 0, 0

	for i := 1; i <= n*n; i++ {
		table[r][c] = i

		nr, nc := r+directions[d][0], c+directions[d][1]
		if nr < 0 || nr >= n || nc < 0 || nc >= n || table[nr][nc] > 0 {
			d = (d + 1) % 4
			nr, nc = r+directions[d][0], c+directions[d][1]
		}
		r, c = nr, nc
	}

	return table
}

func RunTestSpiralMatrixII() {
	runner.InitMetrics("SpiralMatrixII")

	testCases := map[string]struct {
		n      int
		expect [][]int
	}{
		"example-1-n-equals-3": {
			n: 3,
			expect: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		"example-2-n-equals-1": {
			n:      1,
			expect: [][]int{{1}},
		},
		"n-equals-2": {
			n: 2,
			expect: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		"n-equals-4": {
			n: 4,
			expect: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		"n-equals-5": {
			n: 5,
			expect: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
		"edge-case-n-equals-20": {
			n:      20,
			expect: buildSpiralMatrixIIEdgeCaseNEquals20Expect(),
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"n": testCase.n})

		result := runner.ExecCountMetrics(generateMatrix, testCase.n).([][]int)
		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

func buildSpiralMatrixIIEdgeCaseNEquals20Expect() [][]int {
	matrix := make([][]int, 20)
	for i := range matrix {
		matrix[i] = make([]int, 20)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0
	row, col := 0, 0

	for num := 1; num <= 400; num++ {
		matrix[row][col] = num

		nextRow := row + directions[dirIdx][0]
		nextCol := col + directions[dirIdx][1]

		if nextRow < 0 || nextRow >= 20 || nextCol < 0 || nextCol >= 20 || matrix[nextRow][nextCol] != 0 {
			dirIdx = (dirIdx + 1) % 4
			nextRow = row + directions[dirIdx][0]
			nextCol = col + directions[dirIdx][1]
		}

		row, col = nextRow, nextCol
	}

	return matrix
}
