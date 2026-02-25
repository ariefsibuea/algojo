package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("MinimumWindowSubstring", RunTestMinWindowSubstring)
}

/*
 * LeetCode Problem : Minimum Window Substring
 * Topics           : Hash Table, String, Sliding Window
 * Level            : Hard
 * URL              : https://leetcode.com/problems/minimum-window-substring
 * Description      : Given two strings, 's' and 't', return the minimum substring of 's' such that each character in
 * 					't', including duplicates, are included in this substring. By "minimum", I mean the shortest
 * 					substring. If two substrings that satisfy the condition have the same length, the one that comes
 * 					lexicographically first is smaller.
 * Examples         :
 * 					Example 1:
 * 					Input: s = "ADOBECODEBANC", t = "ABC"
 * 					Output: "BANC"
 * 					Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 *
 * 					Example 2:
 * 					Input: s = "a", t = "a"
 * 					Output: "a"
 * 					Explanation: The entire string s is the minimum window.
 *
 * 					Example 3:
 * 					Input: s = "a", t = "aa"
 * 					Output: ""
 * 					Explanation: Both 'a's from t must be included in the window.
 * 					Since the largest window of s only has one 'a', return empty string.
 */

func minWindowSubstring(s string, t string) string {
	requiredChars := 0
	target := [128]int{}
	for _, r := range t {
		if target[r] == 0 {
			requiredChars++
		}
		target[r]++
	}

	matchedChars := 0
	minWindowLength := math.MaxInt

	start := 0
	resultStart := 0

	source := [128]int{}
	for end, r := range s {
		source[r]++

		if target[r] > 0 && target[r] == source[r] {
			matchedChars++
		}

		for requiredChars == matchedChars {
			currentLenght := (end - start) + 1
			if currentLenght < minWindowLength {
				minWindowLength = currentLenght
				resultStart = start
			}

			leftChar := s[start]
			source[leftChar]--

			if target[leftChar] > 0 && target[leftChar] > source[leftChar] {
				matchedChars--
			}

			start++
		}
	}

	if minWindowLength == math.MaxInt {
		return ""
	}

	return s[resultStart : resultStart+minWindowLength]
}

func RunTestMinWindowSubstring() {
	testCases := map[string]struct {
		s      string
		t      string
		expect string
	}{
		"case-1": {
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minWindowSubstring(testCase.s, testCase.t)
		if !cmp.EqualStrings(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
