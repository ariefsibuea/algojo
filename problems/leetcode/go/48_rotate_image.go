package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("RotateImage", RunTestRotateImage)
}

/*
 * Problem	: Rotate Image
 * Topics	: Array, Math, Matrix
 * Level	: Medium
 * URL		: https://leetcode.com/problems/rotate-image/
 *
 * Description:
 *		Given an n x n 2D matrix representing an image, rotate the image by 90 degrees clockwise. The rotation must be
 * 		performed in-place, meaning the input matrix should be modified directly without allocating additional 2D
 * 		matrix space.
 *
 * Constraints:
 *		- n == matrix.length == matrix[i].length
 *		- 1 <= n <= 20
 *		- -1000 <= matrix[i][j] <= 1000
 *
 * Examples:
 *		Example 1:
 *		Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 *		Output: [[7,4,1],[8,5,2],[9,6,3]]
 *		Explanation: After rotating 90 degrees clockwise, the first row becomes the last column, the second row becomes
 * 		the middle column, and the third row becomes the first column.
 *
 *		Example 2:
 *		Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 *		Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 */

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			matrix[i][j], matrix[j][n-1-i] = matrix[j][n-1-i], matrix[i][j]
			matrix[i][j], matrix[n-1-i][n-1-j] = matrix[n-1-i][n-1-j], matrix[i][j]
			matrix[i][j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j]
		}
	}
}

func RunTestRotateImage() {
	runner.InitMetrics("RotateImage")

	testCases := map[string]struct {
		matrix [][]int
		expect [][]int
	}{
		"example-1-3x3-matrix": {
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expect: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		"example-2-4x4-matrix": {
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expect: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		"edge-case-1x1-matrix": {
			matrix: [][]int{
				{1},
			},
			expect: [][]int{
				{1},
			},
		},
		"edge-case-2x2-matrix": {
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expect: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		"matrix-with-negative-numbers": {
			matrix: [][]int{
				{-1, -2},
				{-3, -4},
			},
			expect: [][]int{
				{-3, -1},
				{-4, -2},
			},
		},
		"matrix-with-zeros": {
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expect: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		"5x5-matrix": {
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			expect: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		matrixCopy := make([][]int, len(testCase.matrix))
		for i := range testCase.matrix {
			matrixCopy[i] = make([]int, len(testCase.matrix[i]))
			copy(matrixCopy[i], testCase.matrix[i])
		}

		format.PrintInput(map[string]interface{}{"matrix": matrixCopy})

		runner.ExecCountMetrics(rotate, testCase.matrix)
		if !cmp.EqualSlices(testCase.matrix, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, testCase.matrix)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
