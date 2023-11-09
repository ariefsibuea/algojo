/* Longest Substring Without Repeating Characters
Source		: https://leetcode.com/problems/longest-substring-without-repeating-characters/
Level		: Medium
Description	: Given a string `s`, find the length of the longest `substring` without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Solution Source:
- https://medium.com/outco/how-to-solve-sliding-window-problems-28d67601a66
- https://www.geeksforgeeks.org/window-sliding-technique/
- https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/
*/

package leetcode

import (
	"strings"
)

func (soln Solution) LengthOfLongestSubstring(s string) int {
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
