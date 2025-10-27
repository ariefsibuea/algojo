package main

import (
	"fmt"
	"os"
)

/**
 * Problem 			: First Unique Character in a String
 * Topics           : Hash Table, String, Queue, Counting
 * Level            : Easy
 * URL              : https://leetcode.com/problems/first-unique-character-in-a-string/description/
 * Description      : Given a string `s`, find the first non-repeating character in it and return its index. If it
 * 					does not exist, return -1.
 * Examples         :
 * 					Example 1:
 * 					Input: s = "leetcode"
 * 					Output: 0
 * 					Explanation:
 * 					The character 'l' at index 0 is the first character that does not occur at any other index.
 *
 * 					Example 2:
 * 					Input: s = "loveleetcode"
 * 					Output: 2
 *
 * 					Example 3:
 * 					Input: s = "aabb"
 * 					Output: -1
 */

func firstUniqChar(s string) int {
	// constraint: s consists only lowercase
	chars := [26]int{}
	for _, c := range s {
		chars[c-'a'] += 1
	}

	for i, c := range s {
		if chars[c-'a'] == 1 {
			return i
		}
	}

	return -1
}

func RunTestFirstUniqChar() {
	testCases := map[string]struct {
		s      string
		expect int
	}{
		"case-1": {
			s:      "leetcode",
			expect: 0,
		},
		"case-2": {
			s:      "loveleetcode",
			expect: 2,
		},
		"case-3": {
			s:      "aabb",
			expect: -1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := firstUniqChar(testCase.s)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
