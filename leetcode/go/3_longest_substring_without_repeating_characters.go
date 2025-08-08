package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Longest Substring Without Repeating Characters
Topic            : Hash Table, String, Sliding Window
Level            : Medium
URL              : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
Description      : Given a string s, find the length of the longest substring without duplicate characters.
Examples         :
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
*/

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int) // map of rune to index
	maxLength := 0
	start := 0

	for i, c := range s {
		if _, ok := charMap[c]; ok && charMap[c] >= start {
			start = charMap[c] + 1
		}

		charMap[c] = i
		maxLength = max(maxLength, (i-start)+1)
	}

	return maxLength
}

func max(n, t int) int {
	if n > t {
		return n
	}
	return t
}

func RunTestLongestSubstringWithoutRepeatingCharacters() {
	testCases := map[string]struct {
		s      string
		expect int
	}{
		"case-1": {
			s:      "abcabcbb",
			expect: 3,
		},
		"case-2": {
			s:      "bbbbb",
			expect: 1,
		},
		"case-3": {
			s:      "pwwkew",
			expect: 3,
		},
		"case-4": {
			s:      "abba",
			expect: 2,
		},
	}

	var result int

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result = lengthOfLongestSubstring(testCase.s)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
