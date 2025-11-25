package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Transaction Pair Sum
 * Topics           : Array, Hash Map
 * Level            : Medium
 * Description      : You are given an integer array transactions where transactions[i] represents the amount of the
 * 					  i-th transaction (positive for buy, negative for sell). You are also given an integer target.
 * 					  Return the number of pairs (i, j) such that:
 * 						- 0 <= i < j < len(transactions)
 * 						- transactions[i] + transactions[j] == target
 * 					  Order of the pair doesn’t matter ((i, j) is the same as (j, i); you only count it once because of
 * 					  i < j).
 * Examples         :
 * 					  Example 1:
 * 					  transactions = [2, 7, 11, -2, 4], target = 9
 * 					  Pairs:
 * 					  - (0, 1): 2 + 7 = 9
 * 					  - (2, 3): 11 + (-2) = 9
 * 					  Output: 2
 *
 * 					  Example 2:
 * 					  transactions = [1, 1, 1], target = 2
 * 					  Pairs:
 * 					  (0, 1), (0, 2), (1, 2)
 * 					  Output: 3
 *
 * 					  Example 3:
 * 					  transactions = [3, 5, 7], target = 100
 * 					  No pairs sum to 100.
 * 					  Output: 0
 */

func transactionPairSum_HashMap(transactions []int, target int) int {
	var result = 0
	var freq = map[int]int{}
	var need = 0

	for i := 0; i < len(transactions); i++ {
		need = target - transactions[i]
		if freq[need] > 0 {
			result += freq[need]
		}

		freq[transactions[i]] = freq[transactions[i]] + 1
	}

	return result
}

func transactionPairSum_BruteForce(transactions []int, target int) int {
	var result = 0

	for i := 0; i < len(transactions)-1; i++ {
		for j := i + 1; j < len(transactions); j++ {
			if transactions[i]+transactions[j] == target {
				result += 1
			}
		}
	}

	return result
}

func RunTestTransactionPairSum() {
	testCases := map[string]struct {
		transactions []int
		target       int
		expect       int
	}{
		"case-1": {
			transactions: []int{2, 7, 11, -2, 4},
			target:       9,
			expect:       2,
		},
		"case-2": {
			transactions: []int{1, 1, 1},
			target:       2,
			expect:       3,
		},
		"case-3": {
			transactions: []int{3, 5, 7},
			target:       100,
			expect:       0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := transactionPairSum_HashMap(testCase.transactions, testCase.target)
		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
