package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Climbing Stairs
 * Topic            : Math, Dynamic Programming, Memoization
 * Level            : Easy
 * URL              : https://leetcode.com/problems/climbing-stairs/
 * Description      : You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb
 * 					  1 or 2 steps. In how many distinct ways can you climb to the top?
 * Examples         :
 *         Example 1:
 *         Input: n = 2
 *         Output: 2
 *         Explanation: There are two ways to climb to the top.
 *         1. 1 step + 1 step
 *         2. 2 steps
 *
 *         Example 2:
 *         Input: n = 3
 *         Output: 3
 *         Explanation: There are three ways to climb to the top.
 *         1. 1 step + 1 step + 1 step
 *         2. 1 step + 2 steps
 *         3. 2 steps + 1 step
 */

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	lastOneStep, lastTwoStep := 2, 1
	for i := 3; i <= n; i++ {
		temp := lastOneStep
		lastOneStep = lastOneStep + lastTwoStep
		lastTwoStep = temp
	}

	return lastOneStep
}

func RunTestClimbStairs() {
	testCases := map[string]struct {
		n      int
		expect int
	}{
		"case-1": {
			n:      2,
			expect: 2,
		},
		"case-2": {
			n:      3,
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := climbStairs(testCase.n)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
