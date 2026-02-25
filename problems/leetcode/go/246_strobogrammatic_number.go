package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("StrobogrammaticNumber", RunTestStrobogrammaticNumber)
}

/**
 * Problem 			: Strobogrammatic Numbber
 * Topics           : Hash Table, Two Pointers, String
 * Level            : Easy
 * URL              : https://leetcode.com/problems/strobogrammatic-number
 * Description      : Given a string 'num' representing an integer, determine whether it is a strobogrammatic number.
 * 					Return TRUE if the number is strobogrammatic or FALSE if it is not. A strobogrammatic number
 * 					appears the same when rotated 180 degrees (viewed upside down). For example, “69” is
 * 					strobogrammatic because it looks the same when flipped upside down, while “962” is not.
 * Examples         :
 * 					Example 1:
 * 					Input: str = "808"
 * 					Output: true
 *
 * 					Example 2:
 * 					Input: str = "123"
 * 					Output: false
 *
 * 					Example 3:
 * 					Input: str = "69"
 * 					Output: true
 */

func isStrobogrammatic(num string) bool {
	validNumber := map[byte]byte{
		'0': '0',
		'1': '1',
		'8': '8',
		'6': '9',
		'9': '6',
	}

	start, end := 0, len(num)-1

	for start <= end {
		if val, ok := validNumber[num[start]]; !ok || val != num[end] {
			return false
		}

		start += 1
		end -= 1
	}

	return true
}

func RunTestStrobogrammaticNumber() {
	testCases := map[string]struct {
		num    string
		expect bool
	}{
		"case-1": {
			num:    "808",
			expect: true,
		},
		"case-2": {
			num:    "123",
			expect: false,
		},
		"case-3": {
			num:    "69",
			expect: true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := isStrobogrammatic(testCase.num)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
