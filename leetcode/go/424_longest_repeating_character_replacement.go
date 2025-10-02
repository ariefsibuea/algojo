package main

/**
 * LeetCode Problem : Longest Repeating Character Replacement
 * Topics           : Hash Table, String, Sliding Window
 * Level            : Medium
 * URL              : https://leetcode.com/problems/longest-repeating-character-replacement
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: s = "ABAB", k = 2
 * 					Output: 4
 * 					Explanation: Replace the two 'A's with two 'B's or vice versa.
 *
 * 					Example 2:
 * 					Input: s = "AABABBA", k = 1
 * 					Output: 4
 * 					Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
 * 					The substring "BBBB" has the longest repeating letters, which is 4.
 * 					There may exists other ways to achieve this answer too.
 * Reference		: https://www.hellointerview.com/learn/code/sliding-window/longest-repeating-character-replacement
 */

func characterReplacement(s string, k int) int {
	charCount := map[byte]int{}
	maxFreq := 0   // maximum frequency of any character
	maxLength := 0 // maximum length of the substring

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		charCount[char] += 1
		maxFreq = max(maxFreq, charCount[char])

		// NOTE: In this step we don't decrease the frequency of the maxFreq character. However, we don't have to since
		// it will not affect the final result. Actually, we search the character with frequency > maxFreq. If the
		// maxFreq doesn't change -- even when we have moved far beyond the substring where maxFreq is obtained -- then
		// the final result must be the lenght of substring containing the final maxFreq.
		for (right-left+1)-maxFreq > k {
			prevChar := s[left]
			charCount[prevChar] -= 1
			left += 1
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
