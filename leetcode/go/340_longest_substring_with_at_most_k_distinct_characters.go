package main

import (
	"fmt"
	"os"
)

/**
 * Problem 			: Longest Substring with At Most K Distinct Characters
 * Topics           : Hash Table, String, Sliding Window
 * Level            : Medium
 * URL              : https://www.lintcode.com/problem/386/
 * Description      : Given a string `s` and an integer `k`. Find the length of the longest substring within s that
 * 					contains at most `k` distinct characters.
 * Examples         :
 * 					Example 1:
 * 					Input: S = "eceba" and k = 3
 * 					Output: 4
 * 					Explanation: T = "eceb"
 *
 * 					Example 2:
 * 					Input: S = "WORLD" and k = 4
 * 					Output: 4
 * 					Explanation: T = "WORL" or "ORLD"
 */

// lengthOfLongestSubstringKDistinct provides more efficient solution than SolutionI since we don't need to track
// longest substring and use unnecessary loop. Every time the rule of 'k' is violated, we only need to move the left
// pointer by one position forward instead of shrinking the substring.
//
// Reference: https://algo.monster/liteproblems/340
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	chars := []rune(s)
	charCount := map[rune]int{}
	start := 0

	for _, char := range chars {
		count := charCount[char]
		charCount[char] = count + 1

		if len(charCount) > k {
			remove := chars[start]
			charCount[remove] -= 1
			if charCount[remove] == 0 {
				delete(charCount, remove)
			}
			start += 1
		}
	}

	return len(s) - start
}

func lengthOfLongestSubstringKDistinctSolutionI(s string, k int) int {
	chars := []rune(s)
	charCount := map[rune]int{}
	length, start := 0, 0

	for end, char := range chars {
		count := charCount[char]
		charCount[char] = count + 1

		for len(charCount) > k {
			remove := chars[start]
			charCount[remove] -= 1
			if charCount[remove] == 0 {
				delete(charCount, remove)
			}
			start += 1
		}

		length = max(length, (end-start)+1)
	}

	return length
}

func RunTestLengthOfLongestSubstringKDistinct() {
	testCases := map[string]struct {
		s      string
		k      int
		expect int
	}{
		"case-1": {
			s:      "eceba",
			k:      3,
			expect: 4,
		},
		"case-2": {
			s:      "WORLD",
			k:      4,
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := lengthOfLongestSubstringKDistinct(testCase.s, testCase.k)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
