package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("FibonacciNumber", RunTestFibonacciNumber)
}

/*
 * Problem	: Fibonacci Number
 * Topics	: Math, Dynamic Programming, Recursion, Memoization
 * Level	: Easy
 * URL		: https://leetcode.com/problems/fibonacci-number/
 *
 * Description:
 * 		Compute the nth Fibonacci number, where the sequence is defined by the recurrence relation
 * 		F(n) = F(n-1) + F(n-2) for n > 1, with base values F(0) = 0 and F(1) = 1. Given a non-negative
 * 		integer n, return the corresponding Fibonacci number F(n).
 *
 * Constraints:
 * 		- 0 <= n <= 30
 *
 * Examples:
 * 		Example 1:
 * 		Input: n = 2
 * 		Output: 1
 * 		Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 *
 * 		Example 2:
 * 		Input: n = 3
 * 		Output: 2
 * 		Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
 *
 * 		Example 3:
 * 		Input: n = 4
 * 		Output: 3
 * 		Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 */

func fib(n int) int {
	if n <= 1 {
		return n
	}

	fibseq := make([]int, n+1)
	fibseq[0] = 0
	fibseq[1] = 1

	for i := 2; i <= n; i++ {
		fibseq[i] = fibseq[i-1] + fibseq[i-2]
	}

	return fibseq[n]
}

func RunTestFibonacciNumber() {
	runner.InitMetrics("FibonacciNumber")

	testCases := map[string]struct {
		n      int
		expect int
	}{
		"base-f-0": {
			n:      0,
			expect: 0,
		},
		"base-f-1": {
			n:      1,
			expect: 1,
		},
		"example-2": {
			n:      2,
			expect: 1,
		},
		"example-3": {
			n:      3,
			expect: 2,
		},
		"example-4": {
			n:      4,
			expect: 3,
		},
		"larger-n-10": {
			n:      10,
			expect: 55,
		},
		"max-constraint-30": {
			n:      30,
			expect: 832040,
		},
	}

	var passedCount int

	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"n": tc.n})

		result := runner.ExecCountMetrics(fib, tc.n).(int)
		if !cmp.EqualNumbers(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
