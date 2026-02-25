package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("NumberOfIsland", RunTestNumberOfIsland)
}

/*
 * Problem 			: Number of Island
 * Topics           : Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
 * Level            : Medium
 * URL              : https://leetcode.com/problems/number-of-islands
 * Description      : Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return
 * 					  the number of islands. An island is surrounded by water and is formed by connecting adjacent
 * 					  lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by
 * 					  water.
 * Examples         :
 * 					  Example 1:
 * 					  Input: grid = [
 * 					    ["1","1","1","1","0"],
 * 					    ["1","1","0","1","0"],
 * 					    ["1","1","0","0","0"],
 * 					    ["0","0","0","0","0"]
 * 					  ]
 * 					  Output: 1
 *
 * 					  Example 2:
 * 					  Input: grid = [
 * 					    ["1","1","0","0","0"],
 * 					    ["1","1","0","0","0"],
 * 					    ["0","0","1","0","0"],
 * 					    ["0","0","0","1","1"]
 * 					  ]
 * 					  Output: 3
 */

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var (
		result = 0
		height = len(grid)
		length = len(grid[0])
	)

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'

		for _, direct := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			nextRow, nextCol := i+direct[0], j+direct[1]

			if (0 <= nextRow && nextRow < height) &&
				(0 <= nextCol && nextCol < length) &&
				grid[nextRow][nextCol] == '1' {
				dfs(nextRow, nextCol)
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				result += 1
			}
		}
	}

	return result
}

func RunTestNumberOfIsland() {
	testCases := map[string]struct {
		grid   [][]byte
		expect int
	}{
		"case-1": {
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expect: 1,
		},
		"case-2": {
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := numIslands(testCase.grid)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
