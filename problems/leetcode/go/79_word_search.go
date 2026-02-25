package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("WordSearch", RunTestWordSearch)
}

/**
 * Problem 			: Word Search
 * Topics           : Array, String, Backtracking, Depth-First Search, Matrix
 * Level            : Medium
 * URL              : https://leetcode.com/problems/word-search
 * Description      : Given an m x n grid of characters board and a string 'word', return true if 'word' exists in the
 * 					grid. The 'word' can be constructed from letters of sequentially adjacent cells, where adjacent
 * 					cells are horizontally or vertically neighboring. The same letter cell may not be used more than
 * 					once.
 * Examples         :
 * 					Example 1:
 * 					Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
 * 					Output: true
 *
 * 					Example 2:
 * 					Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
 * 					Output: true
 *
 * 					Example 3:
 * 					Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
 * 					Output: false
 */

func wordSearchExist(board [][]byte, word string) bool {
	var checkBoard func(row, col, wordPos int) bool
	checkBoard = func(row, col, wordPos int) bool {
		if wordPos == len(word)-1 {
			return board[row][col] == word[wordPos]
		}

		if board[row][col] != word[wordPos] {
			return false
		}

		prevVal := board[row][col]
		board[row][col] = 'V' // marked as visited

		nextVisits := [][]int{
			{row - 1, col}, // up
			{row + 1, col}, // bottom
			{row, col - 1}, // left
			{row, col + 1}, // right
		}

		for _, nextVisit := range nextVisits {
			if nextVisit[0] < 0 || nextVisit[0] >= len(board) {
				continue
			}
			if nextVisit[1] < 0 || nextVisit[1] >= len(board[nextVisit[0]]) {
				continue
			}

			if checkBoard(nextVisit[0], nextVisit[1], wordPos+1) {
				return true
			}
		}

		board[row][col] = prevVal // revert value
		return false
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if checkBoard(row, col, 0) {
				return true
			}
		}
	}

	return false
}

func RunTestWordSearch() {
	testCases := map[string]struct {
		board  [][]byte
		word   string
		expect bool
	}{
		"case-1": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "ABCCED",
			expect: true,
		},
		"case-2": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "SEE",
			expect: true,
		},
		"case-3": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "ABCB",
			expect: false,
		},
		"case-4": {
			board: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			word:   "AB",
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := wordSearchExist(testCase.board, testCase.word)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
