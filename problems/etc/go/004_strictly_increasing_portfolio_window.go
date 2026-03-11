package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("StrictlyIncreasingPortfolioWindows", RunTestStrictlyIncreasingPortfolioWindows)
}

/*
 * Problem 			: Strictly Increasing Portfolio Windows
 * Topics           : Array, Two Pointers
 * Level            : Medium
 * Description      : You are given an integer array values where values[i] is a user’s portfolio value at the end of
 * 					  day i (0-indexed). You are also given an integer k.
 * 					  A subarray of length k (i.e., k consecutive elements) is called strictly increasing if:
 * 							values[i] < values[i+1] < ... < values[i+k−1]
 * 					  Return the number of strictly increasing subarrays of length k.
 * Examples         :
 * 					  Example 1:
 * 					  values = [5, 6, 7, 1, 2, 3], k = 3
 * 					  Strictly increasing subarrays of length 3:
 * 					  [5, 6, 7]
 * 					  [1, 2, 3]
 * 					  Output: 2
 *
 * 					  Example 2:
 * 					  values = [3, 3, 4, 5, 2], k = 2
 * 					  Strictly increasing subarrays of length 2:
 * 					  [3, 4]
 * 					  [4, 5]
 * 					  Output: 2
 *
 * 					  Example 3:
 * 					  values = [10, 20, 30], k = 1
 * 					  Every single element forms a strictly increasing subarray of length 1.
 * 					  Output: 3
 */

func strictlyIncreasingPortfolioWindows(values []int, k int) int {
	if len(values) == 0 {
		return 0
	}
	if k == 1 {
		return len(values)
	}

	var result = 0
	var start = 0

	for end := 1; end < len(values); end++ {
		if values[end] <= values[end-1] {
			start = end
		}

		if end-start >= k-1 {
			result += 1
		}
	}

	return result
}

func RunTestStrictlyIncreasingPortfolioWindows() {
	runner.InitMetrics("StrictlyIncreasingPortfolioWindows")

	testCases := map[string]struct {
		values []int
		k      int
		expect int
	}{
		"case-1": {
			values: []int{5, 6, 7, 1, 2, 3},
			k:      3,
			expect: 2,
		},
		"case-2": {
			values: []int{3, 3, 4, 5, 2},
			k:      2,
			expect: 2,
		},
		"case-3": {
			values: []int{4, 3, 2, 1},
			k:      2,
			expect: 0,
		},
		"case-4": {
			values: []int{10, 20, 30},
			k:      1,
			expect: 3,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"values": testCase.values, "k": testCase.k})

		result := runner.ExecCountMetrics(strictlyIncreasingPortfolioWindows, testCase.values, testCase.k).(int)
		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
