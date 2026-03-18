package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SpiralMatrix", RunTestSpiralMatrix)
}

/*
 * Problem	: Spiral Matrix
 * Topics	: Array, Matrix, Simulation
 * Level	: Medium
 * URL		: https://leetcode.com/problems/spiral-matrix/
 *
 * Description:
 *		Given an m x n matrix, return all elements of the matrix in spiral order. The spiral order starts from the
 *		top-left corner and proceeds in a clockwise spiral pattern: first moving right along the top row, then down
 *		along the rightmost column, then left along the bottom row, and finally up along the leftmost column. This
 *		pattern continues, layer by layer, until all elements have been visited.
 *
 * Constraints:
 *		- m == matrix.length
 *		- n == matrix[i].length
 *		- 1 <= m, n <= 10
 *		- -100 <= matrix[i][j] <= 100
 *
 * Examples:
 *		Example 1:
 *		Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 *		Output: [1,2,3,6,9,8,7,4,5]
 *		Explanation: Traverse right (1,2,3), then down (6,9), then left (8,7), then up (4), then right (5).
 *
 *		Example 2:
 *		Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 *		Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *		Explanation: Traverse right (1,2,3,4), then down (8,12), then left (11,10,9), then up (5), then right (6,7).
 */

func spiralOrder(matrix [][]int) []int {
	return spiralOrderSolutions.withIteration(matrix)
}

type spiralOrderSolution struct{}

var spiralOrderSolutions = spiralOrderSolution{}

func (s *spiralOrderSolution) withRecursive(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // bottom
		{0, -1}, // left
		{-1, 0}, // top
	}

	totalCell := m * n
	result := make([]int, 0, totalCell)

	var traverse func(d, r, c int)

	traverse = func(d, r, c int) {
		result = append(result, matrix[r][c])
		matrix[r][c] = 101
		if len(result) == totalCell {
			return
		}

		nr, nc := r+directions[d][0], c+directions[d][1]
		if nr < 0 || nr >= m || nc < 0 || nc >= n || matrix[nr][nc] == 101 {
			d = (d + 1) % 4
			nr, nc = r+directions[d][0], c+directions[d][1]
		}
		traverse(d, nr, nc)
	}

	traverse(0, 0, 0)
	return result
}

func (s *spiralOrderSolution) withIteration(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // bottom
		{0, -1}, // left
		{-1, 0}, // top
	}

	totalCell := m * n
	result := make([]int, 0, totalCell)

	d, r, c := 0, 0, 0

	for i := 0; i < totalCell; i++ {
		result = append(result, matrix[r][c])
		matrix[r][c] = 101

		nr, nc := r+directions[d][0], c+directions[d][1]
		if nr < 0 || nr >= m || nc < 0 || nc >= n || matrix[nr][nc] == 101 {
			d = (d + 1) % 4
			nr, nc = r+directions[d][0], c+directions[d][1]
		}
		r, c = nr, nc
	}

	return result
}

func RunTestSpiralMatrix() {
	runner.InitMetrics("SpiralMatrix")

	testCases := map[string]struct {
		matrix [][]int
		expect []int
	}{
		"example-1-three-by-three": {
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		"example-2-three-by-four": {
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		"single-row": {
			matrix: [][]int{
				{1, 2, 3, 4},
			},
			expect: []int{1, 2, 3, 4},
		},
		"single-column": {
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
			expect: []int{1, 2, 3, 4},
		},
		"one-by-one": {
			matrix: [][]int{
				{1},
			},
			expect: []int{1},
		},
		"empty-matrix": {
			matrix: [][]int{},
			expect: []int{},
		},
		"two-by-two": {
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expect: []int{1, 2, 4, 3},
		},
		"two-by-three": {
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expect: []int{1, 2, 3, 6, 5, 4},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"matrix": testCase.matrix})

		matrixCopy := make([][]int, len(testCase.matrix))
		for i := range testCase.matrix {
			matrixCopy[i] = make([]int, len(testCase.matrix[i]))
			copy(matrixCopy[i], testCase.matrix[i])
		}

		result := runner.ExecCountMetrics(spiralOrder, matrixCopy).([]int)
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
