package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("NumberOfSmoothDescentPeriodsOfAStock", RunTestNumberOfSmoothDescentPeriodsOfAStock)
}

/*
 * Problem 			: Number of Smooth Descent Periods of a Stock
 * Topics           : Array, Math, Dynamic Programming
 * Level            : Medium
 * URL              : https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock
 * Description      : <Description>
 * Examples         : <Examples>
 */

func getDescentPeriods(prices []int) int64 {
	var smoothDescentTotal = int64(0)
	var smoothDescentSubTotal = int64(1)
	var smoothDescentPeriod = int64(1)

	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] != 1 {
			smoothDescentTotal += smoothDescentSubTotal
			smoothDescentSubTotal = 0
			smoothDescentPeriod = 0
		}

		smoothDescentPeriod += 1
		smoothDescentSubTotal += smoothDescentPeriod
	}

	return smoothDescentTotal + smoothDescentSubTotal
}

func RunTestNumberOfSmoothDescentPeriodsOfAStock() {
	testCases := map[string]struct {
		prices []int
		expect int64
	}{
		"case-1": {
			prices: []int{3, 2, 1, 4},
			expect: 7,
		},
		"case-2": {
			prices: []int{8, 6, 7, 7},
			expect: 4,
		},
		"case-3": {
			prices: []int{1},
			expect: 1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := getDescentPeriods(testCase.prices)
		format.PrintInput(map[string]interface{}{"prices": testCase.prices})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
