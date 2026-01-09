package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Two Strings
 * Topics           : Hash Sets, Strings, Frequency Counting
 * Level            : Easy
 * URL              : https://www.hackerrank.com/challenges/two-strings/problem
 * Description      : Given two lowercase strings, determine if they share at least one common character. A shared
 *                    character means there exists a letter that appears in both strings at least once, regardless of
 *                    order or frequency. If such a character exists, return YES; otherwise, return NO. The challenge
 *                    focuses on efficiently detecting any overlap without comparing every substring combination.
 * Examples         : Input: "hello", "world" -> Output: YES (common character 'o')
 *              	  Input: "hi", "world"    -> Output: NO  (no letters overlap)
 */

func twoStrings(s1 string, s2 string) string {
	var hasCharacters [26]bool

	for i := 0; i < len(s1); i++ {
		hasCharacters[s1[i]-'a'] = true
	}

	for i := 0; i < len(s2); i++ {
		if hasCharacters[s2[i]-'a'] {
			return "YES"
		}
	}

	return "NO"
}

func RunTestTwoStrings() {
	testCases := map[string]struct {
		s1     string
		s2     string
		expect string
	}{
		"case-1": {
			s1:     "hello",
			s2:     "world",
			expect: "YES",
		},
		"case-2": {
			s1:     "hi",
			s2:     "world",
			expect: "NO",
		},
		"case-3": {
			s1:     "wouldyoulikefries",
			s2:     "abcabcabcabcabcabc",
			expect: "NO",
		},
		"case-4": {
			s1:     "hackerrankcommunity",
			s2:     "cdecdecdecde",
			expect: "YES",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := twoStrings(testCase.s1, testCase.s2)
		format.PrintInput(map[string]interface{}{"s1": testCase.s1, "s2": testCase.s2})

		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
