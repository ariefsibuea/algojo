package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("LongestPalindromicSubstring", RunTestLongestPalindromicSubstring)
}

/*
 * Problem	: Longest Palindromic Substring
 * Topics	: String, Dynamic Programming
 * Level	: Medium
 * URL		: https://leetcode.com/problems/longest-palindromic-substring/
 *
 * Description:
 * 		Given a string s, find and return the longest substring that reads the same forwards and backwards. A palindrome
 * 		is a sequence of characters that is identical when read from left to right and right to left. The solution
 * 		should identify the longest such contiguous substring within the given input string.
 *
 * Constraints:
 * 		- 1 <= s.length <= 1000
 * 		- s consists of only digits and English letters
 *
 * Examples:
 * 		Example 1:
 * 		Input: s = "babad"
 * 		Output: "bab"
 * 		Explanation: "bab" is a palindrome and appears as a substring. Note that "aba" is also a valid palindrome
 * 		substring of the same length, but the problem accepts any valid answer.
 *
 * 		Example 2:
 * 		Input: s = "cbbd"
 * 		Output: "bb"
 * 		Explanation: "bb" is the longest palindromic substring. It reads the same from both directions.
 */

func longestPalindrome(s string) string {
	return longestPalindromeSolutions.withPalindromeTable(s)
}

type longestPalindromeSolution struct{}

var longestPalindromeSolutions = longestPalindromeSolution{}

func (lps *longestPalindromeSolution) withExpandAroundCenter(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	maxLength := 1
	start := 0

	for i := range s {
		oddLength := expandAroundCenter(i, i)
		evenLength := expandAroundCenter(i, i+1)

		currMaxLength := max(oddLength, evenLength)
		if currMaxLength > maxLength {
			maxLength = currMaxLength
			start = i - (currMaxLength-1)/2
		}
	}

	return s[start : start+maxLength]
}

func (lps *longestPalindromeSolution) withPalindromeTable(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	table := make([][]bool, len(s))
	for i := range table {
		table[i] = make([]bool, len(s))
		table[i][i] = true
	}

	maxLength := 1
	start := 0

	for length := 2; length <= len(s); length++ {
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1

			if s[i] == s[j] {
				if j-i == 1 {
					table[i][j] = true
				} else {
					table[i][j] = table[i+1][j-1]
				}
			}

			if table[i][j] && length > maxLength {
				maxLength = length
				start = i
			}
		}
	}

	return s[start : start+maxLength]
}

func RunTestLongestPalindromicSubstring() {
	runner.InitMetrics("LongestPalindromicSubstring")

	testCases := map[string]struct {
		s      string
		expect string
	}{
		"example-1-multiple-valid-answers": {
			s:      "babad",
			expect: "bab",
		},
		"example-2-even-length-palindrome": {
			s:      "cbbd",
			expect: "bb",
		},
		"single-character": {
			s:      "a",
			expect: "a",
		},
		"entire-string-palindrome": {
			s:      "racecar",
			expect: "racecar",
		},
		"no-palindrome-longer-than-one": {
			s:      "abc",
			expect: "a",
		},
		"even-length-palindrome-full": {
			s:      "abba",
			expect: "abba",
		},
		"all-same-characters": {
			s:      "aaaa",
			expect: "aaaa",
		},
		"palindrome-with-digits": {
			s:      "12321",
			expect: "12321",
		},
		"two-same-characters": {
			s:      "aa",
			expect: "aa",
		},
		"palindrome-at-end": {
			s:      "abacdfgdcaba",
			expect: "aba",
		},
		"mixed-palindromes": {
			s:      "xabayz",
			expect: "aba",
		},
		"longest-in-middle": {
			s:      "xabax",
			expect: "xabax",
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		result := runner.ExecCountMetrics(longestPalindrome, testCase.s).(string)
		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
