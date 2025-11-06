package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Reverse String
 * Topics           : Two Pointers, String
 * Level            : Easy
 * URL              : https://leetcode.com/problems/reverse-string
 * Description      : You are given a character array 's' representing a string. Write a function that reverses the
 * 					array in-place.
 * 						- The reversal must be done by modifying the original array directly.
 * 						- You cannot use extra memory beyond a constant amount O(1).
 * Examples         :
 * 					Example 1:
 * 					Input: s = ["h","e","l","l","o"]
 * 					Output: ["o","l","l","e","h"]
 *
 * 					Example 2:
 * 					Input: s = ["H","a","n","n","a","h"]
 * 					Output: ["h","a","n","n","a","H"]
 */

func reverseString(s []byte) {
	start, end := 0, len(s)-1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		start += 1
		end -= 1
	}
}

func RunTestReverseString() {
	testCases := map[string]struct {
		s      []byte
		expect []byte
	}{
		"case-1": {
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			expect: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		"case-2": {
			s:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expect: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		reverseString(testCase.s)
		if !cmp.EqualSlices(testCase.s, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, testCase.s)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
