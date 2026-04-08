package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("GameOfLife", RunTestGameOfLife)
}

/*
 * Problem	: Game of Life
 * Topics	: Array, Matrix, Simulation
 * Level	: Medium
 * URL		: https://leetcode.com/problems/game-of-life/
 *
 * Description:
 * 		Conway's Game of Life is a cellular automaton played on an m x n grid where each cell is either live (1) or
 * 		dead (0). Each cell evolves based on its eight neighbors (horizontal, vertical, and diagonal) using four rules:
 * 			- live cells with fewer than two neighbors die from under-population
 * 			- live cells with two or three neighbors survive to the next generation
 * 			- live cells with more than three neighbors die from over-population
 * 			- dead cells with exactly three neighbors become alive through reproduction
 * 		All state changes occur simultaneously in one transition.
 *
 * Constraints:
 * 		- m == board.length
 * 		- n == board[i].length
 * 		- 1 <= m, n <= 25
 * 		- board[i][j] is 0 or 1
 *
 * Examples:
 * 		Example 1:
 * 		Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
 * 		Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
 * 		Explanation: The board transitions to the next state. The live cell at position (0,1) dies due to
 * 		under-population (only 1 live neighbor). The dead cell at position (1,0) becomes alive because it has exactly 3
 * 		live neighbors.
 *
 * 		Example 2:
 * 		Input: board = [[1,1],[1,0]]
 * 		Output: [[1,1],[1,1]]
 * 		Explanation: The dead cell at position (1,1) has exactly 3 live neighbors, so it becomes alive through
 * 		reproduction. All other live cells survive with 2-3 neighbors each.
 */

var neighborDirs = [8][2]int{
	{0, -1},  // left
	{0, 1},   // right
	{-1, 0},  // top
	{1, 0},   // bottom
	{-1, -1}, // top-left
	{-1, 1},  // top-right
	{1, -1},  // bottom-left
	{1, 1},   // bottom-right
}

func gameOfLife(board [][]int) {
	const dying, reviving = 2, -1

	rows, cols := len(board), len(board[0])

	for r := range board {
		for c := range board[r] {
			liveNeighbors := 0
			for _, n := range neighborDirs {
				nr, nc := r+n[0], c+n[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if board[nr][nc] >= 1 {
					liveNeighbors++
				}
			}

			switch board[r][c] {
			case 1:
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[r][c] = dying // was alive -> will die
				}
			case 0:
				if liveNeighbors == 3 {
					board[r][c] = reviving // was die -> will alive
				}
			}
		}
	}

	for r := range board {
		for c := range board[r] {
			switch board[r][c] {
			case 2:
				board[r][c] = 0
			case -1:
				board[r][c] = 1
			}
		}
	}
}

func RunTestGameOfLife() {
	runner.InitMetrics("GameOfLife")

	testCases := map[string]struct {
		board  [][]int
		expect [][]int
	}{
		"example-1-glider-pattern": {
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expect: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		"example-2-block-with-dead-cell": {
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			expect: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		"single-live-cell": {
			board: [][]int{
				{1},
			},
			expect: [][]int{
				{0},
			},
		},
		"all-dead-cells": {
			board: [][]int{
				{0, 0},
				{0, 0},
			},
			expect: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		"block-still-life": {
			board: [][]int{
				{1, 1},
				{1, 1},
			},
			expect: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		"blinker-oscillator": {
			board: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			expect: [][]int{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
		},
		"overcrowded-3x3": {
			board: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expect: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		"single-row-pattern": {
			board: [][]int{
				{1, 0, 1, 0},
			},
			expect: [][]int{
				{0, 0, 0, 0},
			},
		},
		"toad-oscillator": {
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
				{0, 0, 0, 0},
			},
			expect: [][]int{
				{0, 0, 1, 0},
				{1, 0, 0, 1},
				{1, 0, 0, 1},
				{0, 1, 0, 0},
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"board": testCase.board})

		boardCopy := make([][]int, len(testCase.board))
		for i := range testCase.board {
			boardCopy[i] = make([]int, len(testCase.board[i]))
			copy(boardCopy[i], testCase.board[i])
		}

		runner.ExecCountMetrics(gameOfLife, testCase.board)

		if !cmp.EqualSlices(testCase.board, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, testCase.board)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
