package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
LeetCode Problem : Best Time to Buy and Sell Stock
Topic            : Array, Dynamic Programming
Level            : Easy
URL              : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
Description      : You are given an array prices where prices[i] is the price of a given stock on the ith day. You want
        to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to
        sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any
        profit, return 0.
Examples         :
        Example 1:
        Input: prices = [7,1,5,3,6,4]
        Output: 5
        Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

        Example 2:
        Input: prices = [7,6,4,3,1]
        Output: 0
        Explanation: In this case, no transactions are done and the max profit = 0.
*/

func maxProfit(prices []int) int {
	idxBuy := 0
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[idxBuy] {
			idxBuy = i
		} else {
			maxProfit = max(maxProfit, prices[i]-prices[idxBuy])
		}
	}

	return maxProfit
}

func RunTestMaxProfit() {
	testCases := map[string]struct {
		prices []int
		expect int
	}{
		"case-1": {
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
		"case-2": {
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := maxProfit(testCase.prices)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
