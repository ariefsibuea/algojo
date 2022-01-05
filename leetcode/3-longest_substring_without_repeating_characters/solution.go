package longestsubstringwithoutrepeatingcharacters

import (
	"strings"
)

/**
 * Problem source: https://leetcode.com/problems/longest-substring-without-repeating-characters/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day VI Sliding Window
 *		Level: Medium
 * Solution source: https://medium.com/outco/how-to-solve-sliding-window-problems-28d67601a66
 * 					https://www.geeksforgeeks.org/window-sliding-technique/
 * 					https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/
**/

// LengthOfLongestSubstring implements window sliding technique to solve the problem.
func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	sub := ""
	for i := range s {
		if strings.Contains(sub, string(s[i])) {
			subs := strings.Split(sub, string(s[i]))
			sub = subs[1]
		}
		sub = sub + string(s[i])
		maxLength = max(maxLength, len(sub))
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
