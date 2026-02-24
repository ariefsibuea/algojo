package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem		: Permutation in String
 * Topics		: Hash Table, Two Pointers, String, Sliding Window
 * Level		: Medium
 * URL			: https://leetcode.com/problems/permutation-in-string
 *
 * Description:
 * 		Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
 * 		In other words, return true if one of s1's permutations is the substring of s2.
 *
 * Constraints:
 * 		- 1 <= s1.length, s2.length <= 10^4.
 * 		- s1 and s2 consist of lowercase English letters.
 *
 * Examples:
 * 		Example 1:
 * 		Input: s1 = "ab", s2 = "eidbaooo"
 * 		Output: true
 * 		Explanation: s2 contains one permutation of s1 ("ba").
 *
 * 		Example 2:
 * 		Input: s1 = "ab", s2 = "eidboaoo"
 * 		Output: false
 */

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	charCount1 := [26]int{}
	for _, c := range s1 {
		charCount1[c-'a']++
	}

	charCount2 := [26]int{}
	start := 0
	for end := 0; end < n2; end++ {
		charCount2[s2[end]-'a']++

		if end-start+1 == n1 {
			if charCount2 == charCount1 {
				return true
			}

			charCount2[s2[start]-'a']--
			start++
		}
	}

	return false
}

func RunTestPermutationInString() {
	runner.InitMetrics("PermutationInString")

	testCases := map[string]struct {
		s1     string
		s2     string
		expect bool
	}{
		"basic-true-permutation-exists": {
			s1:     "ab",
			s2:     "eidbaooo",
			expect: true,
		},
		"basic-false-permutation-does-not-exist": {
			s1:     "ab",
			s2:     "eidboaoo",
			expect: false,
		},
		"s1-longer-than-s2": {
			s1:     "abc",
			s2:     "ab",
			expect: false,
		},
		"s1-and-s2-are-identical": {
			s1:     "abc",
			s2:     "abc",
			expect: true,
		},
		"s1-is-a-permutation-of-s2": {
			s1:     "bac",
			s2:     "abc",
			expect: true,
		},
		"permutation-at-the-beginning": {
			s1:     "ab",
			s2:     "abcooo",
			expect: true,
		},
		"permutation-at-the-end": {
			s1:     "ab",
			s2:     "oooab",
			expect: true,
		},
		"s1-with-duplicates-and-is-a-permutation": {
			s1:     "aab",
			s2:     "baaooo",
			expect: true,
		},
		"true-when-permutation-is-a-substring": {
			s1:     "adc",
			s2:     "dcda",
			expect: true,
		},
		"s1-has-one-char-and-in-s2": {
			s1:     "a",
			s2:     "bca",
			expect: true,
		},
		"s1-has-one-char-and-not-in-s2": {
			s1:     "d",
			s2:     "bca",
			expect: false,
		},
	}

	var passedCount int
	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"s1": testCase.s1, "s2": testCase.s2})

		result := runner.ExecCountMetrics(checkInclusion, testCase.s1, testCase.s2).(bool)
		if !cmp.EqualBooleans(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}
		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
