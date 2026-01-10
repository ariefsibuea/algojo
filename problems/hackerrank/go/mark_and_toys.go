package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Mark and Toys
 * Topics           : Greedy, Sorting, Arrays
 * Level            : Easy
 * URL              : https://www.hackerrank.com/challenges/mark-and-toys/problem
 * Description      : Given a list of toy prices and a budget k, determine the maximum number of toys that can be bought
 *                    without exceeding the budget. The optimal strategy is to sort the prices in ascending order and
 * 					  buy the cheapest toys first until the running total would surpass k.
 * Examples         : Prices: [1 12 5 111 200 1000 10], k=50	-> Output: 4 (buy toys priced 1, 5, 10, 12)
 *              	  Prices: [3 7 2 9 4], k=15					-> Output: 3 (buy toys priced 2, 3, 4)
 */

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	var totalSpent int32 = 0
	var count int32 = 0

	for _, p := range prices {
		if totalSpent+p > k {
			break
		}
		totalSpent += p
		count++
	}

	return count
}

func RunTestMarkAndToys() {
	testCases := map[string]struct {
		prices []int32
		k      int32
		expect int32
	}{
		"case-1": {
			prices: []int32{1, 12, 5, 111, 200, 1000, 10},
			k:      50,
			expect: 4,
		},
		"case-2": {
			prices: []int32{1, 2, 3, 4},
			k:      7,
			expect: 3,
		},
		"case-3": {
			prices: []int32{3, 7, 2, 9, 4},
			k:      15,
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := maximumToys(testCase.prices, testCase.k)
		format.PrintInput(map[string]interface{}{"prices": testCase.prices, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
