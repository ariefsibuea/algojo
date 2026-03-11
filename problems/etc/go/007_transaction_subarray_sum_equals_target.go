package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("TransactionSubarraySumEqualsTarget", RunTestTransactionSubarraySumEqualsTarget)
}

/*
 * Problem 			: Transaction Subarray Sum Equals Target
 * Topics           : Array, Hash Map
 * Level            : Medium
 * Description      : You are given an integer array transactions where transactions[i] represents the amount of the
 * 					  i-th transaction (positive or negative). You are also given an integer target.
 * 					  Return true if there exists a non-empty contiguous subarray whose sum is exactly target,
 * 					  otherwise return false.
 * Examples         :
 * 					  Example 1:
 * 					  transactions = [3, 4, -2, 1, 2, -6], target = 5
 * 					  Possible subarray:
 * 					  [3, 4, -2] = 5
 * 					  Output: true
 *
 * 					  Example 2:
 * 					  transactions = [1, 2, 3, 4], target = 10
 * 					  Subarray [1, 2, 3, 4] sums to 10.
 * 					  Output: true
 *
 * 					  Example 3:
 * 					  transactions = [1, 2, 3], target = 7
 * 					  No contiguous subarray sums to 7.
 * 					  Output: false
 */

func transactionSubarraySumEqualsTarget(transactions []int, target int) bool {
	var prefixSum = transactions[0]
	var prefixSumExist = map[int]bool{transactions[0]: true}

	for i := 1; i < len(transactions); i++ {
		prefixSum += transactions[i]
		if prefixSum == target {
			return true
		}

		need := prefixSum - target
		if prefixSumExist[need] {
			return true
		}

		prefixSumExist[prefixSum] = true
	}
	return false
}

func RunTestTransactionSubarraySumEqualsTarget() {
	runner.InitMetrics("TransactionSubarraySumEqualsTarget")

	testCases := map[string]struct {
		transactions []int
		target       int
		expect       bool
	}{
		"case-1": {
			transactions: []int{3, 4, -2, 1, 2, -6},
			target:       5,
			expect:       true,
		},
		"case-2": {
			transactions: []int{1, 2, 3, 4},
			target:       10,
			expect:       true,
		},
		"case-3": {
			transactions: []int{1, 2, 3},
			target:       7,
			expect:       false,
		},
		"case-4": {
			transactions: []int{1, 2, 3, 0},
			target:       0,
			expect:       true,
		},
		"case-5": {
			transactions: []int{-3, -2},
			target:       -2,
			expect:       true,
		},
		"case-6": {
			transactions: []int{3, 4, -2, 1, 2, -6},
			target:       -3,
			expect:       true,
		},
		"case-7": {
			transactions: []int{1, 2, 3},
			target:       0,
			expect:       false,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"transactions": testCase.transactions, "target": testCase.target})

		result := runner.ExecCountMetrics(transactionSubarraySumEqualsTarget, testCase.transactions, testCase.target).(bool)
		if !cmp.EqualBooleans(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
