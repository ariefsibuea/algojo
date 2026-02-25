package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("HappyNumber", RunTestHappyNumber)
}

/**
 * Problem 			: Happy Number
 * Topics           : Hash Table, Math, Two Pointers
 * Level            : Easy
 * URL              : https://leetcode.com/problems/happy-number
 * Description      : Write an algorithm to determine if a number 'n' is happy. A happy number is a number defined by
 * 					  the following process:
 * 						- Starting with any positive integer, replace the number by the sum of the squares of its
 * 							digits.
 * 						- Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in
 * 							a cycle which does not include 1.
 * 						- Those numbers for which this process ends in 1 are happy.
 * 					  Return true if 'n' is a happy number, and false if not.
 * Examples         :
 * 					  Example 1:
 * 					  Input: n = 19
 * 					  Output: true
 * 					  Explanation:
 * 						1^2 + 9^2 = 82
 * 						8^2 + 2^2 = 68
 * 						6^2 + 8^2 = 100
 * 						1^2 + 0^2 + 0^2 = 1
 *
 * 					  Example 2:
 * 					  Input: n = 2
 * 					  Output: false
 */

func isHappyNumber(n int) bool {
	switch n {
	case 0:
		return false
	case 1:
		return true
	default:
		slow, fast := n, squareOfDigits(n)
		fastStep := 0

		for fast != 1 && fast != slow {
			if fastStep == 2 {
				fastStep = 0
				slow = squareOfDigits(slow)
			} else {
				fast = squareOfDigits(fast)
				fastStep += 1
			}
		}

		return fast == 1
	}
}

func squareOfDigits(num int) int {
	result := 0
	for num > 0 {
		mod := num % 10
		result += mod * mod
		num /= 10
	}
	return result
}

func RunTestHappyNumber() {
	testCases := map[string]struct {
		n      int
		expect bool
	}{
		"case-1": {
			n:      19,
			expect: true,
		},
		"case-2": {
			n:      2,
			expect: false,
		},
		"case-3": {
			n:      1,
			expect: true,
		},
		"case-4": {
			n:      2147483646,
			expect: false,
		},
		"case-5": {
			n:      8,
			expect: false,
		},
		"case-6": {
			n:      7,
			expect: true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := isHappyNumber(testCase.n)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
