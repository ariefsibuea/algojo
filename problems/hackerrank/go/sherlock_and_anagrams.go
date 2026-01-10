package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Sherlock and Anagrams
 * Topics           : Strings, Hash Maps, Combinatorics
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
 * Description      : Given a string, count the number of unordered pairs of substrings that are anagrams of each
 *                    other. Two substrings form a pair when they contain the same multiset of characters regardless of
 *                    their positions in the original string. The task requires examining all substring lengths, using
 *                    frequency signatures to detect equal compositions, and summing the pair counts for every matching
 *                    signature.
 * Examples         : Input: "abba" -> Output: 4 (pairs: "a"-"a", "b"-"b", "ab"-"ba", "abb"-"bba")
 *              	  Input: "abcd" -> Output: 0 (no two substrings share identical character composition)
 */

func sherlockAndAnagrams(s string) int32 {
	n := len(s)

	prefixSum := make([][26]int, n+1)
	for i, c := range s {
		prefixSum[i+1] = prefixSum[i]
		prefixSum[i+1][c-'a'] += 1
	}

	var freqSignature = make(map[[26]int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			// generate signature based on 'prefixSum'
			signature := [26]int{}
			for c := 0; c < 26; c++ {
				signature[c] = prefixSum[j][c] - prefixSum[i][c]
			}
			freqSignature[signature] += 1
		}
	}

	var result int32

	for _, f := range freqSignature {
		result += int32(f * (f - 1) / 2)
	}

	return result
}

func RunTestSherlockAndAnagrams() {
	testCases := map[string]struct {
		s      string
		expect int32
	}{
		"case-1": {
			s:      "abba",
			expect: 4,
		},
		"case-2": {
			s:      "abcd",
			expect: 0,
		},
		"case-3": {
			s:      "ifailuhkqq",
			expect: 3,
		},
		"case-4": {
			s:      "kkkk",
			expect: 10,
		},
		"case-5": {
			s:      "cdcd",
			expect: 5,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := sherlockAndAnagrams(testCase.s)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
