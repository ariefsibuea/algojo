package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("AppendCharactersToStringToMakeSubsequence", RunTestAppendCharactersToString)
}

/**
 * Problem 			: Append Characters to String to Make Subsequence
 * Topics           : Two Pointers, String, Greedy
 * Level            : Medium
 * URL              : https://leetcode.com/problems/append-characters-to-string-to-make-subsequence
 * Description      : You’re given two strings, 's' and 't', made up of lowercase English letters. Your task is to
 * 					determine the minimum number of characters that must be appended to the end of the 's' so that the
 * 					't' becomes a subsequence of the resulting string.
 * Examples         :
 * 					Example 1:
 * 					Input: s = "coaching", t = "coding"
 * 					Output: 4
 * 					Explanation: Append the characters "ding" to the end of s so that s = "coachingding".
 * 					Now, t is a subsequence of s ("coachingding").
 * 					It can be shown that appending any 3 characters to the end of s will never make t a subsequence.
 *
 * 					Example 2:
 * 					Input: s = "abcde", t = "a"
 * 					Output: 0
 * 					Explanation: t is already a subsequence of s ("abcde").
 *
 * 					Example 3:
 * 					Input: s = "z", t = "abcde"
 * 					Output: 5
 * 					Explanation: Append the characters "abcde" to the end of s so that s = "zabcde".
 * 					Now, t is a subsequence of s ("zabcde").
 * 					It can be shown that appending any 4 characters to the end of s will never make t a subsequence.
 */

func appendCharacters(s string, t string) int {
	sourcePtr, targetPtr := 0, 0
	sourceLen := len(s)
	targetLen := len(t)

	for sourcePtr < sourceLen && targetPtr < targetLen {
		if s[sourcePtr] == t[targetPtr] {
			targetPtr += 1
		}
		sourcePtr += 1
	}

	if targetPtr == targetLen {
		return 0
	}

	return targetLen - targetPtr
}

func RunTestAppendCharactersToString() {
	testCases := map[string]struct {
		s      string
		t      string
		expect int
	}{
		"case-1": {
			s:      "coaching",
			t:      "coding",
			expect: 4,
		},
		"case-2": {
			s:      "abcde",
			t:      "a",
			expect: 0,
		},
		"case-3": {
			s:      "z",
			t:      "abcde",
			expect: 5,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := appendCharacters(testCase.s, testCase.t)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
