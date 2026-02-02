package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem          : Minimum Window Subsequence
 * Topics           : String, Dynamic Programming, Sliding Window
 * Level            : Hard
 * URL              : https://leetcode.com/problems/minimum-window-subsequence
 * Description      : Given two strings s1 and s2, find the minimum contiguous substring of s1 such that s2 appears as a
 *                    subsequence within that substring. Characters of s2 must appear in order but not necessarily
 *                    consecutively. If multiple valid substrings of same minimum length exist, return the one with the
 *                    left-most starting index. Return empty string if no valid window exists.
 * Constraints      : - 1 <= s1.length <= 2 * 10^4
 *                    - 1 <= s2.length <= 100
 *                    - s1 and s2 consist of lowercase English letters only
 * Examples         :
 *                    Example 1:
 *                    Input: s1 = "abcdebdde", s2 = "bde"
 *                    Output: "bcde"
 *                    Explanation: Both "bcde" and "bdde" contain s2 as subsequence with length 4, but "bcde" appears
 *                    first so it's the answer. "deb" is not valid because characters must appear in order.
 *
 *                    Example 2:
 *                    Input: s1 = "jmeqksfrsdcmsiwvaovztaqenprpvnbstl", s2 = "u"
 *                    Output: ""
 *                    Explanation: Character 'u' is not present in s1, so no valid window exists.
 */

func minWindow(s1 string, s2 string) string {
	n1, n2 := len(s1), len(s2)
	if n1 < n2 || n2 == 0 {
		return ""
	}

	result := ""
	s2Ptr := 0

	start, end := 0, 0
	for end < n1 {
		if s1[end] == s2[s2Ptr] {
			s2Ptr++
		}

		if s2Ptr >= n2 {
			s2Ptr = n2 - 1
			s1Ptr := end
			tempResult := ""

			for s1Ptr >= start {
				if s1[s1Ptr] == s2[s2Ptr] {
					s2Ptr--
				}

				if s2Ptr < 0 {
					tempResult = s1[s1Ptr : end+1]
					break
				}
				s1Ptr--
			}

			if len(tempResult) < len(result) || result == "" {
				result = tempResult
			}

			s2Ptr = 0
			end = s1Ptr
		}

		end++
	}

	return result
}

func RunTestMinimumWindowSubsequence() {
	testCases := map[string]struct {
		s1     string
		s2     string
		expect string
	}{
		"basic-case": {
			s1:     "abcdebdde",
			s2:     "bde",
			expect: "bcde",
		},
		"no-valid-window": {
			s1:     "jmeqksfrsdcmsiwvaovztaqenprpvnbstl",
			s2:     "u",
			expect: "",
		},
		"exact-match": {
			s1:     "abc",
			s2:     "abc",
			expect: "abc",
		},
		"single-char-match": {
			s1:     "a",
			s2:     "a",
			expect: "a",
		},
		"single-char-no-match": {
			s1:     "a",
			s2:     "b",
			expect: "",
		},
		"subsequence-at-start": {
			s1:     "abcde",
			s2:     "abc",
			expect: "abc",
		},
		"subsequence-at-end": {
			s1:     "abcde",
			s2:     "cde",
			expect: "cde",
		},
		"consecutive-chars": {
			s1:     "aaabbb",
			s2:     "ab",
			expect: "ab",
		},
		"s2-longer-than-s1": {
			s1:     "abc",
			s2:     "abcd",
			expect: "",
		},
		"multiple-windows-leftmost": {
			s1:     "geeksforgeeks",
			s2:     "eksrg",
			expect: "eksforg",
		},
		"repeated-pattern": {
			s1:     "aaaaaaaaaaa",
			s2:     "aa",
			expect: "aa",
		},
		"empty-s2": {
			s1:     "abc",
			s2:     "",
			expect: "",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minWindow(testCase.s1, testCase.s2)
		format.PrintInput(map[string]interface{}{"s1": testCase.s1, "s2": testCase.s2})

		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
