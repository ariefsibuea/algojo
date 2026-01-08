package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem      	: Balanced Brackets
 * Topics       	: Stacks, Strings, Bracket Matching
 * Level        	: Medium
 * URL          	: https://www.hackerrank.com/challenges/balanced-brackets/problem
 * Description  	: For each string composed only of (), {}, and [] characters, determine whether every opening
 * 					  bracket is closed by the same type in the correct order. Return YES if the string is balanced;
 * 					  otherwise, return NO.
 * Examples     	: Input: {[()]}        -> YES
 *              	  Input: {[(])}        -> NO
 *              	  Input: {{[[(())]]}}  -> YES
 */

const (
	BalanceBrackets   = "YES"
	ImbalanceBrackets = "NO"
)

func isBalancedBrackets(s string) string {
	var bracketPair = map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	var bracketStack = make([]rune, 0, len(s))

	for _, c := range s {
		if openBracket, ok := bracketPair[c]; ok {
			if len(bracketStack) == 0 || openBracket != bracketStack[len(bracketStack)-1] {
				return ImbalanceBrackets
			}
			bracketStack[len(bracketStack)-1] = 0
			bracketStack = bracketStack[:len(bracketStack)-1]
		} else {
			bracketStack = append(bracketStack, c)
		}
	}

	if len(bracketStack) > 0 {
		return ImbalanceBrackets
	}
	return BalanceBrackets
}

func RunTestBalancedBrackets() {
	testCases := map[string]struct {
		s      string
		expect string
	}{
		"case-1": {
			s:      "{[()]}",
			expect: "YES",
		},
		"case-2": {
			s:      "{[(])}",
			expect: "NO",
		},
		"case-3": {
			s:      "{{[[(())]]}}",
			expect: "YES",
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := isBalancedBrackets(testCase.s)
		format.PrintInput(map[string]interface{}{"s": testCase.s})

		if !cmp.EqualStrings(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
