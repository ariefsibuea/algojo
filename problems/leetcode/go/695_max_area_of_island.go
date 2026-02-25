package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("MaxAreaOfIsland", RunTestMaxAreaOfIsland)
}

/*
 * Problem 			: Max Area of Island
 * Topics           : Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
 * Level            : Medium
 * URL              : https://leetcode.com/problems/max-area-of-island
 * Description      : You are given an m x n binary matrix grid. An island is a group of 1's (representing land)
 *                    connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid
 *                    are surrounded by water. The area of an island is the number of cells with a value 1 in the
 *                    island. Return the maximum area of an island in grid. If there is no island, return 0.
 * Constraints      :
 *                    - m == grid.length
 *                    - n == grid[i].length
 *                    - 1 <= m, n <= 50
 *                    - grid[i][j] is either 0 or 1.
 * Examples         :
 *                    Example 1:
 *                    Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 *                                   [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *                                   [0,1,1,0,1,0,0,0,0,0,0,0,0],
 *                                   [0,1,0,0,1,1,0,0,1,0,1,0,0],
 *                                   [0,1,0,0,1,1,0,0,1,1,1,0,0],
 *                                   [0,0,0,0,0,0,0,0,0,0,1,0,0],
 *                                   [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *                                   [0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *                    Output: 6
 *                    Explanation: The answer is not 11, because the island must be connected 4-directionally.
 *
 *                    Example 2:
 *                    Input: grid = [[0,0,0,0,0,0,0,0]]
 *                    Output: 0
 */

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		grid[r][c] = 0
		area := 1

		for _, d := range directions {
			r2, c2 := r+d[0], c+d[1]
			if (0 <= r2 && r2 < rows) && (0 <= c2 && c2 < cols) && grid[r2][c2] == 1 {
				area += dfs(r2, c2)
			}
		}

		return area
	}

	maxArea := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				maxArea = max(maxArea, dfs(r, c))
			}
		}
	}

	return maxArea
}

func RunTestMaxAreaOfIsland() {
	// TODO: complete the test function!
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
