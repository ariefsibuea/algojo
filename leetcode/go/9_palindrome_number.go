package main

import (
	"fmt"
	"os"
)

/**
 * Problem 			: Palindrome Number
 * Topics           : Math
 * Level            : Easy
 * URL              : https://leetcode.com/problems/palindrome-number
 * Description      : <Description>
 * Examples         : <Examples>
 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	_x := 0
	divResult := x

	for divResult > 0 {
		mod := divResult % 10
		_x = (_x * 10) + mod
		divResult = divResult / 10

	}

	return _x == x
}

func RunTestIsPalindrome() {
	testCases := map[string]struct {
		x      int
		expect bool
	}{
		"case-1": {
			x:      121,
			expect: true,
		},
		"case-2": {
			x:      -121,
			expect: false,
		},
		"case-3": {
			x:      10,
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := isPalindrome(testCase.x)
		if !EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
