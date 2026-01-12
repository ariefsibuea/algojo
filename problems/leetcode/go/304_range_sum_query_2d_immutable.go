package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Range Sum Query 2D - Immutable
 * Topics           : Array, Matrix, Prefix Sum, Design
 * Level            : Medium
 * URL              : https://leetcode.com/problems/range-sum-query-2d-immutable
 * Description      : Given an m x n integer matrix, design an immutable data structure that supports answering sum
 * 					  queries over any sub-rectangle. The constructor receives the full matrix and can precompute
 * 					  auxiliary data, while the SumRegion(row1, col1, row2, col2) function returns the sum of all matrix
 * 					  elements inside the rectangle whose top-left corner is (row1, col1) and bottom-right corner is
 * 					  (row2, col2). All indices are zero-based and you may assume the parameters always form a valid
 * 					  rectangle within the original matrix.
 * Examples         :
 * 					  Example 1:
 * 					  Input: ["NumMatrix","sumRegion","sumRegion","sumRegion"], [[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],
 * 						[4,1,0,1,7],[1,0,3,0,5]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
 * 					  Output: [null, 8, 11, 12]
 * 					  Explanation: NumMatrix is initialized with the 5x5 grid; querying rows 2..4 & cols 1..3 sums to 8,
 * 					  rows 1..2 & cols 1..2 sums to 11, and rows 1..2 & cols 2..4 sums to 12.
 *
 * 					  Example 2:
 * 					  Input: ["NumMatrix","sumRegion"], [[[1],[2]], [0,0,0,0]]
 * 					  Output: [null, 1]
 * 					  Explanation: With a single-column matrix, querying the entire matrix returns the only element.
 */

type NumMatrix struct {
	prefixSum [][]int
}

func NumMatrixConstructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])

	prefixSum := make([][]int, row+1)
	prefixSum[0] = make([]int, col+1)

	for i := 1; i <= row; i++ {
		prefixSum[i] = make([]int, col+1)

		for j := 1; j <= col; j++ {
			prefixSum[i][j] = matrix[i-1][j-1] +
				(prefixSum[i][j-1] - prefixSum[i-1][j-1]) +
				prefixSum[i-1][j]
		}
	}

	return NumMatrix{prefixSum: prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return (this.prefixSum[row2+1][col2+1] - this.prefixSum[row1][col2+1]) -
		(this.prefixSum[row2+1][col1] - this.prefixSum[row1][col1])
}

func RunTestRangeSumQuery2DImmutable() {
	testCases := map[string]struct{}{}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		fmt.Println(testCase)
		// result := twoSum(testCase.nums, testCase.target)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	os.Exit(1)
		// }
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
