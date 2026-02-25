package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("ValidPalindrome", RunTestValidPalindrome)
}

/**
 * LeetCode 		: Valid Palindrome
 * Topic            : Two Pointers, String
 * Level            : Easy
 * URL              : https://leetcode.com/problems/valid-palindrome/
 * Description      : A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
 * 					removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric
 * 					characters include letters and numbers. Given a string s, return true if it is a palindrome, or
 * 					false otherwise.
 * Examples         :
 * 					Example 1:
 * 					Input: s = "A man, a plan, a canal: Panama"
 * 					Output: true
 * 					Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 * 					Example 2:
 * 					Input: s = "race a car"
 * 					Output: false
 * 					Explanation: "raceacar" is not a palindrome.
 *
 */

func isValidPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start <= end {
		for start < end && !isAlphanumericASCII(s[start]) {
			start += 1
		}
		for start < end && !isAlphanumericASCII(s[end]) {
			end -= 1
		}

		if toLowerCaseASCII(s[start]) != toLowerCaseASCII(s[end]) {
			return false
		}

		start += 1
		end -= 1
	}

	return true
}

func isAlphanumericASCII(c byte) bool {
	switch {
	// numbers
	case '0' <= c && c <= '9':
		return true
	// uppercase alphabet
	case 'A' <= c && c <= 'Z':
		return true
	// lowercase alphabet
	case 'a' <= c && c <= 'z':
		return true
	default:
		return false
	}
}

func toLowerCaseASCII(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func RunTestValidPalindrome() {
	testCases := map[string]struct {
		s      string
		expect bool
	}{
		"case-1": {
			s:      "A man, a plan, a canal: Panama",
			expect: true,
		},
		"case-2": {
			s:      "race a car",
			expect: false,
		},
		"case-3": {
			s:      " ",
			expect: true,
		},
		"case-4": {
			s:      "  a  ",
			expect: true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := isValidPalindrome(testCase.s)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
