package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Count Covered Buildings
 * Topics           : Array, Hash Table, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/count-covered-buildings
 * Description      : <Description>
 * Examples         : <Examples>
 */

func countCoveredBuildings(n int, buildings [][]int) int {
	minMaxRow := make(map[int][]int) // key is column number
	minMaxCol := make(map[int][]int) // key is row number

	for _, b := range buildings {
		x, y := b[0], b[1]

		if _, ok := minMaxRow[y]; !ok {
			minMaxRow[y] = []int{n + 1, 0}
		}
		minMaxRow[y][0] = min(minMaxRow[y][0], x)
		minMaxRow[y][1] = max(minMaxRow[y][1], x)

		if _, ok := minMaxCol[x]; !ok {
			minMaxCol[x] = []int{n + 1, 0}
		}
		minMaxCol[x][0] = min(minMaxCol[x][0], y)
		minMaxCol[x][1] = max(minMaxCol[x][1], y)
	}

	var result int

	for _, b := range buildings {
		x, y := b[0], b[1]

		if (minMaxRow[y][0] < x && x < minMaxRow[y][1]) &&
			(minMaxCol[x][0] < y && y < minMaxCol[x][1]) {
			result += 1
		}
	}

	return result
}

func RunTestCountCoveredBuildings() {
	testCases := map[string]struct {
		n         int
		buildings [][]int
		expect    int
	}{
		"case-1": {
			n: 3,
			buildings: [][]int{
				{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3},
			},
			expect: 1,
		},
		"case-2": {
			n: 3,
			buildings: [][]int{
				{1, 1}, {1, 2}, {2, 1}, {2, 2},
			},
			expect: 0,
		},
		"case-3": {
			n: 5,
			buildings: [][]int{
				{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3},
			},
			expect: 1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := countCoveredBuildings(testCase.n, testCase.buildings)
		format.PrintInput(map[string]interface{}{"n": testCase.n, "buildings": testCase.buildings})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
