package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("LongestRepeatingCharacterReplacement", RunTestCharacterReplacement)
}

/*
 * Problem	: Longest Repeating Character Replacement
 * Topics	: Hash Table, String, Sliding Window
 * Level	: Medium
 * URL		: https://leetcode.com/problems/longest-repeating-character-replacement
 *
 * Description:
 * 		You are given a string s and an integer k. You can choose any character of the string and change it to any
 * 		other uppercase English character. You can perform this operation at most k times. Return the length of the
 * 		longest substring containing the same letter you can get after performing the above operations.
 *
 * Constraints:
 * 		- 1 <= s.length <= 10^5
 * 		- s consists of only uppercase English letters
 * 		- 0 <= k <= s.length
 *
 * Examples:
 * 		Example 1:
 * 		Input: s = "ABAB", k = 2
 * 		Output: 4
 * 		Explanation: Replace the two 'A's with two 'B's or vice versa.
 *
 * 		Example 2:
 * 		Input: s = "AABABBA", k = 1
 * 		Output: 4
 * 		Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA". The substring "BBBB" has the
 * 		longest repeating letters, which is 4. There may exists other ways to achieve this answer too.
 */

func characterReplacement(s string, k int) int {
	charCount := map[byte]int{}
	maxFreq := 0   // maximum frequency of any character
	maxLength := 0 // maximum length of the substring

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		charCount[char] += 1
		maxFreq = max(maxFreq, charCount[char])

		// NOTE: In this step we don't decrease the frequency of the maxFreq character. However, we don't have to since
		// it will not affect the final result. Actually, we search the character with frequency > maxFreq. If the
		// maxFreq doesn't change -- even when we have moved far beyond the substring where maxFreq is obtained -- then
		// the final result must be the lenght of substring containing the final maxFreq.
		for (right-left+1)-maxFreq > k {
			prevChar := s[left]
			charCount[prevChar] -= 1
			left += 1
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func RunTestCharacterReplacement() {
	runner.InitMetrics("LongestRepeatingCharacterReplacement")

	testCases := map[string]struct {
		s      string
		k      int
		expect int
	}{
		"example-1": {
			s:      "ABAB",
			k:      2,
			expect: 4,
		},
		"example-2": {
			s:      "AABABBA",
			k:      1,
			expect: 4,
		},
		"single-char": {
			s:      "A",
			k:      0,
			expect: 1,
		},
		"all-same": {
			s:      "AAAA",
			k:      2,
			expect: 4,
		},
		"k-equals-length": {
			s:      "ABCD",
			k:      4,
			expect: 4,
		},
		"k-zero-no-replace": {
			s:      "AABABBA",
			k:      0,
			expect: 2,
		},
		"all-different-k-one": {
			s:      "ABC",
			k:      1,
			expect: 2,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"s": tc.s, "k": tc.k})

		result := runner.ExecCountMetrics(characterReplacement, tc.s, tc.k).(int)
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
