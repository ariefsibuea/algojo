/* Maximal Square
Source		: https://leetcode.com/problems/maximal-square/
Level		: Medium
Description	: Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and
			return its area.

Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4

Example 2:
Input: matrix = [["0","1"],["1","0"]]
Output: 1

Example 3:
Input: matrix = [["0"]]
Output: 0
*/

package leetcode

func (soln Solution) MaximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])

	var max int
	var prev, curr []int

	curr = make([]int, rows)
	for i := rows - 1; i >= 0; i-- {
		prev, curr = curr, make([]int, cols)
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				continue
			}

			switch {
			case i == rows-1, j == cols-1:
				curr[j] = 1
			default:
				curr[j] = 1 + min(curr[j+1], prev[j], prev[j+1])
			}

			if max < curr[j] {
				max = curr[j]
			}
		}
	}

	return max * max
}

func min(right, down, diagonal int) int {
	min := right
	if min > down {
		min = down
	}
	if min > diagonal {
		min = diagonal
	}
	return min
}
