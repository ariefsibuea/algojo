package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ValidParentheses", RunTestValidParentheses)
}

/*
 * Problem	: Valid Parentheses
 * Topics	: String, Stack
 * Level	: Easy
 * URL		: https://leetcode.com/problems/valid-parentheses/
 *
 * Description:
 * 		Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the
 * 		input string is valid. An input string is valid if:
 * 		1) Open brackets must be closed by the same type of brackets,
 * 		2) Open brackets must be closed in the correct order,
 * 		3) Every close bracket has a corresponding open bracket of the same type.
 *
 * Constraints:
 * 		- 1 <= s.length <= 10^4
 * 		- s consists of parentheses only '()[]{}'.
 *
 * Examples:
 * 		Example 1:
 * 		Input: s = "()"
 * 		Output: true
 *
 * 		Example 2:
 * 		Input: s = "()[]{}"
 * 		Output: true
 *
 * 		Example 3:
 * 		Input: s = "(]"
 * 		Output: false
 *
 * 		Example 4:
 * 		Input: s = "([])"
 * 		Output: true
 */

func isValidParentheses(s string) bool {
	brackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []rune{}
	for _, c := range s {
		if bracket, exists := brackets[c]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != bracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

func RunTestValidParentheses() {
	runner.InitMetrics("ValidParentheses")

	testCases := map[string]struct {
		s      string
		expect bool
	}{
		"example-1-simple": {
			s:      "()",
			expect: true,
		},
		"example-2-mixed": {
			s:      "()[]{}",
			expect: true,
		},
		"example-3-invalid": {
			s:      "(]",
			expect: false,
		},
		"example-4-nested": {
			s:      "([])",
			expect: true,
		},
		"example-5-cross-nested": {
			s:      "([)]",
			expect: false,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := runner.ExecCountMetrics(isValidParentheses, tc.s).(bool)

		format.PrintInput(map[string]interface{}{"s": tc.s})

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
