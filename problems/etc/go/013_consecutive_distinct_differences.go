package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ConsecutiveDistinctDifferences", RunTestConsecutiveDistinctDifferences)
}

/*
 * Problem	: Consecutive Distinct Differences
 * Topics	: Array, Hash Table
 * Level	: Easy
 * URL		: N/A
 *
 * Description:
 * 		Given an array A of N integers, count the number of distinct absolute differences between consecutive elements.
 * 		For each adjacent pair (A[i], A[i+1]), compute |A[i] - A[i+1]| and return how many unique difference values
 * 		result.
 *
 * Constraints:
 * 		- N is an integer within the range [2, 100,000]
 * 		- Each element of A is an integer within the range [-1,000,000,000, 1,000,000,000]
 *
 * Examples:
 * 		Example 1:
 * 		Input: A = [2, 5, 2, 7, 4]
 * 		Output: 2
 * 		Explanation: Consecutive pairs produce differences 3, 3, 5, 3.
 * 		Distinct values are {3, 5}.
 *
 * 		Example 2:
 * 		Input: A = [0, 1, -2, 3, -4, 5]
 * 		Output: 5
 * 		Explanation: Differences are 1, 3, 5, 7, 9 — all distinct.
 *
 * 		Example 3:
 * 		Input: A = [0, 0]
 * 		Output: 1
 * 		Explanation: The single difference is 0.
 */

func distinctDifferences(nums []int) int {
	diffs := make(map[int]bool)

	numDiff := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for i := 1; i < len(nums); i++ {
		diffs[numDiff(nums[i], nums[i-1])] = true
	}

	return len(diffs)
}

func RunTestConsecutiveDistinctDifferences() {
	runner.InitMetrics("ConsecutiveDistinctDifferences")

	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"example-1": {
			nums:   []int{2, 5, 2, 7, 4},
			expect: 2,
		},
		"example-2": {
			nums:   []int{0, 1, -2, 3, -4, 5},
			expect: 5,
		},
		"minimum-array": {
			nums:   []int{0, 0},
			expect: 1,
		},
		"constant-difference": {
			nums:   []int{1, 3, 5, 7},
			expect: 1,
		},
		"unique-differences": {
			nums:   []int{1, 2, 4, 7, 11},
			expect: 4,
		},
		"large-range": {
			nums:   []int{-1000000000, 1000000000},
			expect: 1,
		},
	}

	var passedCount int

	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"nums": tc.nums})

		result := runner.ExecCountMetrics(distinctDifferences, tc.nums).(int)
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
