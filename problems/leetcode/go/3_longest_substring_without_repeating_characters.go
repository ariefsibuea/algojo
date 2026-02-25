package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("LongestSubstringWithoutRepeatingCharacters", RunTestLongestSubstringWithoutRepeatingCharacters)
}

/*
 * Problem	: Longest Substring Without Repeating Characters
 * Topics	: Hash Table, String, Sliding Window
 * Level	: Medium
 * URL		: https://leetcode.com/problems/longest-substring-without-repeating-characters/
 *
 * Description:
 * 		Given a string s, find the length of the longest substring without duplicate characters.
 *
 * Constraints:
 * 		- 0 <= s.length <= 5 * 10^4
 * 		- s consists of English letters, digits, symbols and spaces.
 *
 * Examples:
 * 		Example 1:
 * 		Input: s = "abcabcbb"
 * 		Output: 3
 * 		Explanation: The answer is "abc", with the length of 3.
 *
 * 		Example 2:
 * 		Input: s = "bbbbb"
 * 		Output: 1
 * 		Explanation: The answer is "b", with the length of 1.
 *
 * 		Example 3:
 * 		Input: s = "pwwkew"
 * 		Output: 3
 * 		Explanation: The answer is "wke", with the length of 3.
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

func RunTestLongestSubstringWithoutRepeatingCharacters() {
	runner.InitMetrics("LongestSubstringWithoutRepeatingCharacters")

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

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"s": tc.s})

		result := runner.ExecCountMetrics(lengthOfLongestSubstring, tc.s).(int)
		if !cmp.EqualNumbers(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
