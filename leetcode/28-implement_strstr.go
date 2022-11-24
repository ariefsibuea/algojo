/* Find the Index of the First Occurrence in a String
Source		: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
Level		: Medium
Description	: Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or
			-1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/

package leetcode

import "strings"

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	startIndex := 0
	result := -1
	for {
		if len(haystack) <= len(needle) && haystack != needle {
			result = -1
			break
		}

		if strings.HasPrefix(haystack, needle) {
			result = startIndex
			break
		}

		haystack = haystack[1:]
		startIndex = startIndex + 1
	}

	return result
}
