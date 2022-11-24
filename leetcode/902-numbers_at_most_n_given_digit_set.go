/* Reorder List
Source		: https://leetcode.com/problems/numbers-at-most-n-given-digit-set/description/
Level		: Hard
Description	: Given an array of digits which is sorted in non-decreasing order. You can write numbers using each
			digits[i] as many times as we want. For example, if digits = ['1','3','5'], we may write numbers such as
			'13', '551', and '1351315'. Return the number of positive integers that can be generated that are less than
			or equal to a given integer n.

Example 1:
Input: digits = ["1","3","5","7"], n = 100
Output: 20
Explanation:
The 20 numbers that can be written are:
1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.

Example 2:
Input: digits = ["1","4","9"], n = 1000000000
Output: 29523
Explanation:
We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
In total, this is 29523 integers that can be written using the digits array.

Solution source: https://www.youtube.com/watch?v=nqCM8opotjU
*/

package leetcode

import (
	"math"
	"sort"
	"strconv"
)

func AtMostNGivenDigitSet(digits []string, n int) int {
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})

	numstr := strconv.Itoa(n)
	count := 0
	for i := 1; i <= len(numstr)-1; i++ {
		count += powInt(len(digits), i)
	}

	i := 0
	for i < len(numstr) {
		j := 0
		for j < len(digits) && digits[j][0] < numstr[i] {
			count += powInt(len(digits), len(numstr)-1-i)
			j++
		}
		if j == len(digits) || digits[j][0] > numstr[i] {
			break
		}
		i++
	}

	if i == len(numstr) {
		count++
	}

	return count
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
