package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * Problem 			: Reverse Words in a String
 * Topics           : Two Pointers, String
 * Level            : Medium
 * URL              : https://leetcode.com/problems/reverse-words-in-a-string
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: s = "the sky is blue"
 * 					Output: "blue is sky the"
 *
 * 					Example 2:
 * 					Input: s = "  hello world  "
 * 					Output: "world hello"
 * 					Explanation: Your reversed string should not contain leading or trailing spaces.
 *
 * 					Example 3:
 * 					Input: s = "a good   example"
 * 					Output: "example good a"
 * 					Explanation: You need to reduce multiple spaces between two words to a single space in the
 * 					reversed string.
 */

func reverseWordsS1(s string) string {
	result := []rune{}
	runes := []rune(s)

	right := len(s) - 1
	left := 0

	for right >= 0 {
		if runes[right] != ' ' {
			left = right - 1
			for left >= 0 && runes[left] != ' ' {
				left -= 1
			}

			if len(result) > 0 {
				result = append(result, ' ')
			}

			result = append(result, runes[left+1:right+1]...)
			right = left
		} else {
			right -= 1
		}
	}

	return string(result)
}

func reverseWordsS2(s string) string {
	words := strings.Fields(s)

	left, right := 0, len(words)-1

	for left < right {
		words[left], words[right] = words[right], words[left]
		right -= 1
		left += 1
	}

	return strings.Join(words, " ")
}

func RunTestReverseWordsInAString() {
	testCases := map[string]struct {
		s      string
		expect string
	}{
		"case-1": {
			s:      "the sky is blue",
			expect: "blue is sky the",
		},
		"case-2": {
			s:      "  hello world  ",
			expect: "world hello",
		},
		"case-3": {
			s:      "a good   example",
			expect: "example good a",
		},
		"case-4": {
			s:      "EPY2giL",
			expect: "EPY2giL",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := reverseWordsS2(testCase.s)
		if !EqualStrings(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
