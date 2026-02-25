package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("ValidWordAbbreviation", RunTestValidWordAbbreviation)
}

/**
 * Problem 			: Valid Word Abbreviation
 * Topics           : Two Pointers, String
 * Level            : Easy
 * URL              : https://leetcode.com/problems/valid-word-abbreviation
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: word = "innovation", abbr = "in5ion"
 * 					Output: true
 * 					Explanation: By replacing "5" in abbreviation in "in5ion" with "novat", we form the given word
 * 					"innovation".
 *
 * 					Example 2:
 * 					Input: word = "mindset", abbr = "min3et"
 * 					Output: false
 * 					Explanation: "min3et" is not a correct abbreaviation for "mindset", as replacing "3" with "ds"
 * 					does not reconstruct the original word correctly.
 *
 * 					Example 3:
 * 					Input: word = "leadership", abbr = "lead04ip"
 * 					Output: false
 * 					Explanation: "lead04ip" is not a correct abbreviation for "mindset", as it has a leading 0.
 */

func validWordAbbreviation(word string, abbr string) bool {
	if len(abbr) > len(word) {
		return false
	}
	if abbr == word {
		return true
	}

	wordLength := len(word)
	abbrLength := len(abbr)

	wordIndex := 0
	abbrIndex := 0

	skip := 0
	var currentChar byte

	for wordIndex < wordLength && abbrIndex < abbrLength {
		currentChar = abbr[abbrIndex]

		if '0' <= currentChar && currentChar <= '9' {
			if currentChar == '0' && skip == 0 {
				return false
			}

			digit, err := strconv.Atoi(string(currentChar))
			if err != nil {
				return false
			}

			skip = (skip * 10) + digit
		} else {
			wordIndex += skip
			skip = 0

			if wordIndex >= wordLength || word[wordIndex] != currentChar {
				return false
			}

			wordIndex += 1
		}

		abbrIndex += 1
	}

	return wordIndex+skip == wordLength && abbrIndex == abbrLength
}

func RunTestValidWordAbbreviation() {
	testCases := map[string]struct {
		word   string
		abbr   string
		expect bool
	}{
		"case-1": {
			word:   "innovation",
			abbr:   "in5ion",
			expect: true,
		},
		"case-2": {
			word:   "mindset",
			abbr:   "min3et",
			expect: false,
		},
		"case-3": {
			word:   "leadership",
			abbr:   "lead04ip",
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := validWordAbbreviation(testCase.word, testCase.abbr)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
