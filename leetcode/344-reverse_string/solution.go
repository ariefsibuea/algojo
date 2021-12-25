package reversestring

/**
 * Problem source: https://leetcode.com/problems/reverse-string/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day IV Two Pointers
 *		Level: Easy
 * Solution source:
**/

// ReverseString implements two pointers technique to reverse string
func ReverseString(s []byte) {
	start, end := 0, len(s)-1
	for start < end {
		temp := s[end]
		s[end] = s[start]
		s[start] = temp
		start, end = start+1, end-1
	}
}
