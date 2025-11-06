package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Valid Palindrome II
 * Topics           : Two Pointers, String, Greedy
 * Level            : Easy
 * URL              : https://leetcode.com/problems/valid-palindrome-ii
 * Description      : Write a function that takes a string as input and checks whether it can be a valid palindrome by
 * 					removing at most one character from it.
 * Examples         :
 * 					Example 1:
 * 					Input: s = "aba"
 * 					Output: true
 *
 * 					Example 2:
 * 					Input: s = "abca"
 * 					Output: true
 * 					Explanation: You could delete the character 'c'.
 *
 * 					Example 3:
 * 					Input: s = "abc"
 * 					Output: false
 */

func validPalindromeII(s string) bool {
	start, end := 0, len(s)-1

	isPalindrome := func(start, end int) bool {
		for start <= end {
			if s[start] != s[end] {
				return false
			}

			start += 1
			end -= 1
		}
		return true
	}

	for start <= end {
		if s[start] != s[end] {
			return isPalindrome(start+1, end) || isPalindrome(start, end-1)
		}

		start += 1
		end -= 1
	}

	return true
}

func RunTestValidPalindromeII() {
	testCases := map[string]struct {
		s      string
		expect bool
	}{
		"case-1": {
			s:      "aba",
			expect: true,
		},
		"case-2": {
			s:      "abca",
			expect: true,
		},
		"case-3": {
			s:      "abc",
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := validPalindromeII(testCase.s)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
