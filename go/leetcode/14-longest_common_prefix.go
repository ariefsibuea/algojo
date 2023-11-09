/* Longest Common Prefix
Source		: https://leetcode.com/problems/longest-common-prefix/
Level		: Easy
Description	: Find the longest common prefix string amongst an array of strings. If there is no common prefix, return an
			empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

package leetcode

import "strings"

func (soln Solution) LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(prefix) > len(strs[i]) {
			prefix = prefix[0:len(strs[i])]
		}

		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
		}
	}

	return prefix
}
