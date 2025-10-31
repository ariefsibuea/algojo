package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Proleb	: Ransom Note
 * Topic			: Hash Table, String, Counting
 * Level			: Easy
 * URL				: https://leetcode.com/problems/ransom-note
 * Description		: Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using
 * 					  the letters from magazine and false otherwise. Each letter in magazine can only be used once in
 * 					  ransomNote.
 * Examples			:
 *					  Input: ransomNote = "a", magazine = "b"
 *					  Output: false
 *					  Explanation: magazine doesn't contain letter 'a'
 *
 *					  Input: ransomNote = "aa", magazine = "ab"
 *					  Output: false
 *					  Explanation: magazine contains only one 'a', but ransomNote needs two
 *
 *					  Input: ransomNote = "aa", magazine = "aab"
 *					  Output: true
 *					  Explanation: magazine contains two 'a's and one 'b', sufficient for ransomNote
 */

func canConstructRansomNote(ransomNote string, magazine string) bool {
	charMaps := make(map[byte]int)

	for i := range magazine {
		charMaps[magazine[i]] += 1
	}

	for i := range ransomNote {
		count, exist := charMaps[ransomNote[i]]
		if exist && count > 0 {
			charMaps[ransomNote[i]] -= 1
		} else {
			return false
		}
	}

	return true
}

func RunTestCanConstructRansomNote() {
	testCases := map[string]struct {
		ransomNote string
		magazine   string
		expect     bool
	}{
		"case-1": {
			ransomNote: "a",
			magazine:   "b",
			expect:     false,
		},
		"case-2": {
			ransomNote: "aa",
			magazine:   "ab",
			expect:     false,
		},
		"case-3": {
			ransomNote: "aa",
			magazine:   "aab",
			expect:     true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := canConstructRansomNote(testCase.ransomNote, testCase.magazine)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
