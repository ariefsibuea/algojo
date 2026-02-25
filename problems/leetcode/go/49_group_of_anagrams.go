package main

import (
	"fmt"
	"strings"

	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("GroupAnagrams", RunTestGroupAnagrams)
}

/*
 * Problem	: Group Anagrams
 * Topics	: Array, Hash Table, String, Sorting
 * Level	: Medium
 * URL		: https://leetcode.com/problems/group-anagrams/
 *
 * Description:
 * 		You are given an array of strings strs. Your task is to group all the anagrams together and
 * 		return them as a list of lists. Anagrams are words that contain the same characters but in
 * 		different orders (for example, "eat" and "tea" are anagrams).
 *
 * Constraints:
 * 		- 1 <= strs.length <= 10^4
 * 		- 0 <= strs[i].length <= 100
 * 		- strs[i] consists of lowercase English letters.
 *
 * Examples:
 * 		Example 1:
 * 		Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * 		Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * 		Explanation:
 * 		- There is no string in strs that can be rearranged to form "bat".
 * 		- The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
 * 		- The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each
 * 		other.
 *
 * 		Example 2:
 * 		Input: strs = [""]
 * 		Output: [[""]]
 *
 * 		Example 3:
 * 		Input: strs = ["a"]
 * 		Output: [["a"]]
 */

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		signature := generateSignature(str)
		groups[signature] = append(groups[signature], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func generateSignature(str string) string {
	alphabets := make([]rune, 26)

	for _, c := range str {
		alphabets[c-'a'] += 1
	}

	var signatureBuilder strings.Builder
	for i, count := range alphabets {
		if count > 0 {
			signatureBuilder.WriteRune('a' + rune(i))
			signatureBuilder.WriteString(fmt.Sprintf("%d", count))
		}
	}

	return signatureBuilder.String()
}

func RunTestGroupAnagrams() {
	runner.InitMetrics("GroupAnagrams")

	testCases := map[string]struct {
		strs   []string
		expect [][]string
	}{
		"example-1-basic": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{
				{"eat", "tea", "ate"},
				{"bat"},
				{"tan", "nat"},
			},
		},
		"example-2-empty-string": {
			strs: []string{""},
			expect: [][]string{
				{""},
			},
		},
		"example-3-single-char": {
			strs: []string{"a"},
			expect: [][]string{
				{"a"},
			},
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"strs": tc.strs})

		result := runner.ExecCountMetrics(groupAnagrams, tc.strs).([][]string)
		if !compareAnagramGroups(tc.expect, result) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

func compareAnagramGroups(expected, actual [][]string) bool {
	if len(expected) != len(actual) {
		return false
	}

	used := make([]bool, len(actual))

	for _, expGroup := range expected {
		found := false
		for i, actGroup := range actual {
			if used[i] {
				continue
			}
			if sameGroup(expGroup, actGroup) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func sameGroup(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	countA := make(map[string]int)
	for _, s := range a {
		countA[s]++
	}

	for _, s := range b {
		countA[s]--
		if countA[s] < 0 {
			return false
		}
	}

	for _, c := range countA {
		if c != 0 {
			return false
		}
	}

	return true
}
