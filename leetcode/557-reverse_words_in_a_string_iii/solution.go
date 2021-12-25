package reversewordsinastringiii

/**
 * Problem source: https://leetcode.com/problems/reverse-words-in-a-string-iii/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day IV Two Pointers
 *		Level: Easy
 * Solution source:
**/

import "strings"

func ReverseWords(s string) string {
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
