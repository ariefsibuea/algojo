/* Reverse String
Source		: https://leetcode.com/problems/reverse-string/
Level		: Easy
Description	: Write a function that reverses a string. The input string is given as an array of characters s. You must
			do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

package leetcode

// ReverseString implements two pointers technique to reverse string
func (soln Solution) ReverseString(s []byte) {
	start, end := 0, len(s)-1
	for start < end {
		temp := s[end]
		s[end] = s[start]
		s[start] = temp
		start, end = start+1, end-1
	}
}
