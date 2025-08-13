package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Valid Parentheses
Topic            : String, Stack
Level            : Easy
URL              : https://leetcode.com/problems/valid-parentheses
Description      : Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the
        input string is valid. An input string is valid if: 1) Open brackets must be closed by the same type of brackets,
        2) Open brackets must be closed in the correct order, 3) Every close bracket has a corresponding open bracket of
        the same type.
Examples         :
        Example 1:
        Input: s = "()"
        Output: true

        Example 2:
        Input: s = "()[]{}"
        Output: true

        Example 3:
        Input: s = "(]"
        Output: false

        Example 4:
        Input: s = "([])"
        Output: true
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

func RunTestIsValidParentheses() {
	testCases := map[string]struct {
		s      string
		expect bool
	}{
		"case-1": {
			s:      "()",
			expect: true,
		},
		"case-2": {
			s:      "()[]{}",
			expect: true,
		},
		"case-3": {
			s:      "(]",
			expect: false,
		},
		"case-4": {
			s:      "([])",
			expect: true,
		},
		"case-5": {
			s:      "([)]",
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := isValidParentheses(testCase.s)
		if !EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
