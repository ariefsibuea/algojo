/* Reverse Words in a String III
Source		: https://leetcode.com/problems/reverse-words-in-a-string-iii/
Level		: Easy
Description	: Given a string s, reverse the order of characters in each word within a sentence while still preserving
			whitespace and initial word order.

Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "God Ding"
Output: "doG gniD"
*/

package leetcode

import "strings"

func (soln Solution) ReverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		word := make([]byte, len(words[i]))
		start, end := 0, len(words[i])-1
		for start < end {
			word[start] = words[i][end]
			word[end] = words[i][start]
			start, end = start+1, end-1
		}
		if start == end {
			word[start] = words[i][start]
		}
		words[i] = string(word)
	}

	return strings.Join(words, " ")
}
