/* Add Binary
Source		: https://leetcode.com/problems/add-binary/
Level		: Easy
Description	: Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
*/

package leetcode

import (
	"strconv"
	"strings"
)

func (soln Solution) AddBinary(a string, b string) string {
	ra := []rune(a)
	rb := []rune(b)

	indexA, indexB := len(ra)-1, len(rb)-1
	remainder, binA, binB := 0, 0, 0
	result := make([]string, 0)
	for indexA >= 0 || indexB >= 0 {
		binA, binB = 0, 0
		if indexA >= 0 {
			binA, _ = strconv.Atoi(string(ra[indexA]))
			indexA--
		}
		if indexB >= 0 {
			binB, _ = strconv.Atoi(string(rb[indexB]))
			indexB--
		}

		bin := (binA + binB + remainder) % 2
		remainder = (binA + binB + remainder) / 2
		result = append([]string{strconv.Itoa(bin)}, result...)
	}

	if remainder > 0 {
		result = append([]string{strconv.Itoa(remainder)}, result...)
	}

	return strings.Join(result, "")
}
