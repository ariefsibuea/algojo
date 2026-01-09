package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Ransom Note
 * Topics           : Hash Maps, Strings, Frequency Counting
 * Level            : Easy
 * URL              : https://www.hackerrank.com/challenges/ctci-ransom-note/problem
 * Description      : Determine whether the words needed to write a ransom note can be cut out from a magazine. Each
 *                    magazine word can only be used once, so the note is possible only when every word in the note
 *                    exists in the magazine with at least the same frequency. Output "Yes" when the note can be formed
 *                    and "No" otherwise, taking care to handle duplicate words and differing counts.
 * Examples         : Magazine: "give me one grand today night", Note: "give one grand today" -> Output: Yes
 *              	  Magazine: "two times three is not four", Note: "two times two is four"  -> Output: No
 */

func checkMagazine(magazine []string, note []string) string {
	wordCount := make(map[string]int)

	for _, w := range magazine {
		wordCount[w]++
	}

	for _, w := range note {
		if wordCount[w] == 0 {
			// fmt.Printf("No\n")
			return "No"
		}
		wordCount[w]--
	}

	// fmt.Printf("Yes\n")
	return "Yes"
}

func RunTestRansomNote() {
	testCases := map[string]struct {
		magazine []string
		note     []string
		expect   string
	}{
		"case-1": {
			magazine: []string{
				"give",
				"me",
				"one",
				"grand",
				"today",
				"night",
			},
			note: []string{
				"give",
				"one",
				"grand",
				"today",
			},
			expect: "Yes",
		},
		"case-2": {
			magazine: []string{
				"two",
				"times",
				"three",
				"is",
				"not",
				"four",
			},
			note: []string{
				"two",
				"times",
				"two",
				"is",
				"four",
			},
			expect: "No",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := checkMagazine(testCase.magazine, testCase.note)
		format.PrintInput(map[string]interface{}{"magazine": testCase.magazine, "note": testCase.note})

		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
