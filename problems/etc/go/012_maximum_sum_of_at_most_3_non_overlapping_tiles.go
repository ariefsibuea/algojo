package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("MaximumSumOfAtMost3NonOverlappingTiles", RunTestMaximumSumOfAtMost3NonOverlappingTiles)
}

/*
 * Problem	: Maximum Sum of At Most 3 Non-Overlapping Tiles
 * Topics	: Array, Dynamic Programming
 * Level	: Medium
 * URL		:
 *
 * Description:
 * 		Given an array A of N integers, you may place up to 3 tiles where each tile covers exactly 2 consecutive
 * 		elements. Tiles cannot overlap, cannot extend beyond the array, and each element belongs to at most one
 * 		tile. Determine the maximum possible sum of elements covered by such a placement.
 *
 * Constraints:
 * 		- N is an integer within [2, 100,000]
 * 		- Each element of A is an integer within [0, 1,000,000]
 *
 * Examples:
 * 		Example 1:
 * 		Input: A = [2, 3, 5, 2, 3, 4, 6, 4, 1]
 * 		Output: 25
 * 		Explanation: Tiles placed at (3,5), (3,4), and (6,4) give the maximum sum.
 *
 * 		Example 2:
 * 		Input: A = [1, 5, 3, 2, 6, 6, 10, 4, 7, 2, 1]
 * 		Output: 35
 * 		Explanation: One optimal placement uses tiles at (5,3), (6,10), and (4,7).
 *
 * 		Example 3:
 * 		Input: A = [1, 2, 3, 3, 2]
 * 		Output: 10
 * 		Explanation: Only 2 tiles fit because the array is too small for a third. Tiles at (2,3) and (3,2).
 *
 * 		Example 4:
 * 		Input: A = [5, 10, 3]
 * 		Output: 15
 * 		Explanation: Only 1 tile fits.
 */

func maxSumTiles(nums []int) int {
	maxSum := make([][]int, len(nums)+1) // 1-based index array to simulate position
	for i := range maxSum {
		maxSum[i] = make([]int, 4) // manual set to 4 to represent maxximum number of tiles from 0..3
	}

	for i := 2; i <= len(nums); i++ { // 1-based index to simulate postion
		tile := nums[i-2] + nums[i-1]

		// skip j = 0 since zero tile means sum = 0
		for j := 1; j <= 3; j++ { // represent number of tiles
			// skip current tile, take the maximum sum until 1 position behind
			maxSum[i][j] = maxSum[i-1][j]
			// count current tile, take the maximum sum until 2 position behind for number of tiles maximum at j-1
			// and add itu to current tile
			sumWithTile := maxSum[i-2][j-1] + tile

			maxSum[i][j] = max(maxSum[i][j], sumWithTile)
		}
	}

	best := 0
	for i := 0; i < 4; i++ {
		best = max(best, maxSum[len(nums)][i])
	}

	return best
}

func RunTestMaximumSumOfAtMost3NonOverlappingTiles() {
	runner.InitMetrics("MaximumSumOfAtMost3NonOverlappingTiles")

	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"example-1": {
			nums:   []int{2, 3, 5, 2, 3, 4, 6, 4, 1},
			expect: 25,
		},
		"example-2": {
			nums:   []int{1, 5, 3, 2, 6, 6, 10, 4, 7, 2, 1},
			expect: 35,
		},
		"example-3": {
			nums:   []int{1, 2, 3, 3, 2},
			expect: 10,
		},
		"example-4": {
			nums:   []int{5, 10, 3},
			expect: 15,
		},
		"edge-min-n": {
			nums:   []int{100, 200},
			expect: 300,
		},
		"edge-all-zero": {
			nums:   []int{0, 0, 0, 0, 0, 0, 0, 0},
			expect: 0,
		},
		"max-capacity": {
			nums:   []int{1, 2, 3, 4, 5, 6},
			expect: 21,
		},
		"single-tile-best": {
			nums:   []int{10, 20, 30},
			expect: 50,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"nums": tc.nums})

		result := runner.ExecCountMetrics(maxSumTiles, tc.nums).(int)
		if !cmp.EqualNumbers(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
