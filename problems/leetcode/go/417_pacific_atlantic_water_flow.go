package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("PacificAtlanticWaterFlow", RunTestPacificAtlanticWaterFlow)
}

/*
 * Problem 			: Pacific Atlantic Water Flow
 * Topics           : Array, Depth-First Search, Breadth-First Search, Matrix
 * Level            : Medium
 * URL              : https://leetcode.com/problems/pacific-atlantic-water-flow
 * Description      : There is an m x n rectangular island that borders both the Pacific Ocean and the Atlantic Ocean.
 *                    The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the
 *                    island's right and bottom edges.
 *
 *                    The island is partitioned into a grid of square cells. You are given an m x n integer matrix
 *                    heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
 *
 *                    The island receives a lot of rain, and the rain water can flow to neighboring cells directly
 *                    north, south, east, and west if the neighboring cell's height is less than or equal to the
 *                    current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.
 *
 *                    Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water
 *                    can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
 * Constraints      :
 *                    - m == heights.length
 *                    - n == heights[i].length
 *                    - 1 <= m, n <= 200
 *                    - 0 <= heights[i][j] <= 10^5
 * Examples         :
 *                    Example 1:
 *                    Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
 *                    Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
 *
 *                    Example 2:
 *                    Input: heights = [[2,1],[1,2]]
 *                    Output: [[0,0],[0,1],[1,0],[1,1]]
 */

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])

	visitedPacific := make([][]bool, rows)
	visitedAtlantic := make([][]bool, rows)

	for r := 0; r < rows; r++ {
		visitedPacific[r] = make([]bool, cols)
		visitedAtlantic[r] = make([]bool, cols)
	}

	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true

		for _, d := range directions {
			r2, c2 := r+d[0], c+d[1]

			if (0 <= r2 && r2 < rows) &&
				(0 <= c2 && c2 < cols) &&
				(!visited[r2][c2] && heights[r2][c2] >= heights[r][c]) {

				dfs(r2, c2, visited)
			}
		}
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, visitedPacific)
		dfs(r, cols-1, visitedAtlantic)
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, visitedPacific)
		dfs(rows-1, c, visitedAtlantic)
	}

	var result = make([][]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if visitedPacific[r][c] && visitedAtlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

func RunTestPacificAtlanticWaterFlow() {
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
