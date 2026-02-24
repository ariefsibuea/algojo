package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: Palindrome Number
 * Topics	: Math
 * Level	: Easy
 * URL		: https://leetcode.com/problems/palindrome-number/
 *
 * Description:
 * 		Given an integer x, you need to determine whether it is a palindrome. A palindrome integer reads
 * 		the same backward as forward. Return true if x is a palindrome, and false otherwise.
 *
 * Constraints:
 * 		- -2^31 <= x <= 2^31 - 1
 *
 * Examples:
 * 		Example 1:
 * 		Input: 121
 * 		Output: true
 *
 * 		Example 2:
 * 		Input: -121
 * 		Output: false
 * 		Explanation: From left to right it reads -121, but from right to left it becomes 121-.
 *
 * 		Example 3:
 * 		Input: 10
 * 		Output: false
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

func RunTestPalindromeNumber() {
	runner.InitMetrics("PalindromeNumber")

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

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"x": tc.x})

		result := runner.ExecCountMetrics(isPalindrome, tc.x).(bool)
		if !cmp.EqualBooleans(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
