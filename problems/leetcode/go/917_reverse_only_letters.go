package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Reverse Only Letters
 * Topics           : Two Pointers, String
 * Level            : Easy
 * URL              : https://leetcode.com/problems/reverse-only-letters
 * Description      : <Description>
 * Examples         : <Examples>
 */

func reverseOnlyLetters(s string) string {
	inRune := []rune(s)
	outRune := make([]rune, len(inRune))

	// set non-alphabet character first to reserve slot
	for i := 0; i < len(inRune); i++ {
		if !unicode.IsLetter(inRune[i]) {
			outRune[i] = inRune[i]
		}
	}

	// read the input from end -> start, put alphabet to empty slot
	indexIn := len(inRune) - 1
	indexOut := 0
	for indexIn >= 0 {
		if !unicode.IsLetter(inRune[indexIn]) {
			indexIn--
			continue
		}
		if outRune[indexOut] != 0 {
			indexOut++
			continue
		}

		outRune[indexOut] = inRune[indexIn]
		indexOut++
		indexIn--
	}

	// return result
	return string(outRune)
}

func RunTestReverseOnlyLetters() {
	testCases := map[string]struct {
		s      string
		expect string
	}{
		"case-1": {
			s:      "abc-DEFG-hIjKl",
			expect: "lKj-IhGF-EDcba",
		},
		"case-2": {
			s:      "aBCD*_eFGhIJK!@Lmn",
			expect: "nmLK*_JIhGFeD!@CBa",
		},
		"case-3": {
			s:      "Test1ng-Leet=code-Q!",
			expect: "Qedo1ct-eeLg=ntse-T!",
		},
		"case-4": {
			s:      "-",
			expect: "-",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := reverseOnlyLetters(testCase.s)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
