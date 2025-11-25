package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Maximum Trades in a Time Window
 * Topics           : Array, Two Pointers
 * Level            : Medium
 * Description      : You are given an array of trades for a single stock. Each trade is represented as an integer
 * 					  time[i], the timestamp of the trade in seconds (you can assume trades for this problem are
 * 					  already sorted by time in non-decreasing order).
 * 					  You are also given an integer window, representing a number of seconds.
 * 					  You need to return the maximum number of trades that occur within any contiguous interval of
 * 					  length window seconds.
 * 					  Formally, you want the maximum size of a set of indices {i, i+1, ..., j} such that:
 * 							time[j]−time[i]≤window
 * Examples         :
 * 					  Example 1:
 * 					  time = [1, 2, 3, 10, 11, 12], window = 2
 * 					  Possible windows:
 * 					  - [1, 2, 3] → length 3 trades, time[2] - time[0] = 2
 * 					  - [10, 11, 12] → length 3, time[5] - time[3] = 2
 * 					  Maximum trades in any window of length 2 seconds is 3.
 * 					  Output: 3
 *
 * 					  Example 2:
 * 					  time = [5, 6, 100], window = 1
 * 					  Windows:
 * 					  - [5, 6] → time[1] - time[0] = 1, length = 2
 * 					  - [6] → length = 1
 * 					  - [100] → length = 1
 * 					  Maximum is 2.
 * 					  Output: 2
 *
 * 					  Example 3:
 * 					  time = [10, 20, 30], window = 0
 * 					  Each trade alone fits in a window of length 0 (no two different times can be within 0 seconds
 * 					  unless timestamps are equal).
 * 					  Maximum is 1.
 * 					  Output: 1
 */

func maximumTradesInATimeWindow(times []int, window int) int {
	// If no time defined, return 0 means no trade
	if len(times) == 0 {
		return 0
	}

	var maxTrades = 1

	// Check max trades using two pointers
	// Update max trades each time we find times[end]-times[start] <= window
	start := 0
	for end := 1; end < len(times); end++ {
		if times[end]-times[start] > window {
			// Move start by 1 index just to maintain length of max trades
			start += 1
		} else {
			maxTrades = max(maxTrades, end-start+1)
		}
	}

	return maxTrades
}

func RunTestMaximumTradesInATimeWindow() {
	testCases := map[string]struct {
		times  []int
		window int
		expect int
	}{
		"case-1": {
			times:  []int{1, 2, 3, 10, 11, 12},
			window: 2,
			expect: 3,
		},
		"case-2": {
			times:  []int{5, 6, 100},
			window: 1,
			expect: 2,
		},
		"case-3": {
			times:  []int{10, 20, 30},
			window: 0,
			expect: 1,
		},
		"case-4": {
			times:  []int{1, 10, 20, 30},
			window: 5,
			expect: 1,
		},
		"case-5": {
			times:  []int{1, 1, 1},
			window: 0,
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := maximumTradesInATimeWindow(testCase.times, testCase.window)
		format.PrintInput(map[string]interface{}{"times": testCase.times, "window": testCase.window})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
