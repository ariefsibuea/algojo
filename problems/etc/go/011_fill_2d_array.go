package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("Fill2DArray", RunTestFill2DArray)
}

/*
 * Problem	: Fill 2D Array
 * Topics	: Array, Simulation
 * Level	: Medium
 * URL		: -
 *
 * Description:
 *		Fill a 2D array with monotonically increasing numbers in spiral order, starting from the
 *		top-left corner. Numbers increment only when turning corners, creating concentric layers
 *		of cells with the same value.
 *
 * Constraints:
 *		-
 *
 * Examples:
 *		Example 1:
 *		Input: width = 5, height = 2
 *		Output: [[1 1 1 1 1], [3 3 3 3 2]]
 *
 *		Example 2:
 *		Input: width = 3, height = 3
 *		Output: [[1 1 1], [4 5 2], [3 3 2]]
 */

func fill2DArray(width, height int) [][]int {
	return fill2DArraySolutions.withIterative(width, height)
}

type fill2DArraySolution struct{}

var fill2DArraySolutions = fill2DArraySolution{}

func (s *fill2DArraySolution) withRecursive(width, height int) [][]int {
	table := make([][]int, height)
	for i := range table {
		table[i] = make([]int, width)
	}

	directions := [][2]int{
		{1, 0},  // right
		{0, 1},  // bottom
		{-1, 0}, // left
		{0, -1}, // top
	}
	var fillTable func(direction, r, c, num int) bool

	fillTable = func(d, c, r, num int) bool {
		if c < 0 || c >= width || r < 0 || r >= height {
			return false
		}
		if table[r][c] > 0 {
			return false
		}

		table[r][c] = num
		nc, nr := c+directions[d][0], r+directions[d][1]
		if !fillTable(d, nc, nr, num) {
			d = (d + 1) % 4
			nc, nr = c+directions[d][0], r+directions[d][1]
			fillTable(d, nc, nr, num+1)
		}

		return true
	}

	fillTable(0, 0, 0, 1)
	return table
}

func (s *fill2DArraySolution) withIterative(width, height int) [][]int {
	table := make([][]int, height)
	for i := range table {
		table[i] = make([]int, width)
	}

	directions := [][2]int{
		{1, 0},  // right
		{0, 1},  // bottom
		{-1, 0}, // left
		{0, -1}, // top
	}

	col, row := 0, 0
	hasFilled := func(c, r int) bool {
		return table[r][c] > 0
	}
	outOfTable := func(c, r int) bool {
		return c < 0 || c >= width || r < 0 || r >= height
	}

	num := 1
	dir := 0
	fillCell := width * height

	for filledCell := 0; filledCell < fillCell; filledCell++ {
		table[row][col] = num

		newCol, newRow := col+directions[dir][0], row+directions[dir][1]
		if outOfTable(newCol, newRow) || hasFilled(newCol, newRow) {
			dir = (dir + 1) % 4
			newCol, newRow = col+directions[dir][0], row+directions[dir][1]
			num++
		}
		col, row = newCol, newRow
	}

	return table
}

func RunTestFill2DArray() {
	runner.InitMetrics("Fill2DArray")

	testCases := map[string]struct {
		width  int
		height int
		expect [][]int
	}{
		"example-1-width-5-height-2": {
			width:  5,
			height: 2,
			expect: [][]int{
				{1, 1, 1, 1, 1},
				{3, 3, 3, 3, 2},
			},
		},
		"example-2-width-3-height-3": {
			width:  3,
			height: 3,
			expect: [][]int{
				{1, 1, 1},
				{4, 5, 2},
				{3, 3, 2},
			},
		},
		"single-cell": {
			width:  1,
			height: 1,
			expect: [][]int{{1}},
		},
		"single-row": {
			width:  4,
			height: 1,
			expect: [][]int{{1, 1, 1, 1}},
		},
		"single-column": {
			width:  1,
			height: 4,
			expect: [][]int{{1}, {2}, {2}, {2}},
		},
		"width-4-height-4": {
			width:  4,
			height: 4,
			expect: [][]int{
				{1, 1, 1, 1},
				{4, 5, 5, 2},
				{4, 7, 6, 2},
				{3, 3, 3, 2},
			},
		},
		"width-2-height-2": {
			width:  2,
			height: 2,
			expect: [][]int{{1, 1}, {3, 2}},
		},
		"width-5-height-5": {
			width:  5,
			height: 5,
			expect: [][]int{
				{1, 1, 1, 1, 1},
				{4, 5, 5, 5, 2},
				{4, 8, 9, 6, 2},
				{4, 7, 7, 6, 2},
				{3, 3, 3, 3, 2},
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"width": testCase.width, "height": testCase.height})

		equal2DSlices := func(a, b [][]int) bool {
			if len(a) != len(b) {
				return false
			}
			for i := range a {
				if !cmp.EqualSlices(a[i], b[i]) {
					return false
				}
			}
			return true
		}

		result := runner.ExecCountMetrics(fill2DArray, testCase.width, testCase.height).([][]int)
		if !equal2DSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
