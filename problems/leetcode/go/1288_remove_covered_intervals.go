package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: Remove Covered Intervals
 * Topics	: Array, Sorting
 * Level	: Medium
 * URL		: https://leetcode.com/problems/remove-covered-intervals
 *
 * Description:
 * 		Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri),
 * 		remove all intervals that are covered by another interval in the list.
 * 		The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.
 * 		Return the number of remaining intervals.
 *
 * Constraints:
 * 		- 1 <= intervals.length <= 1000
 * 		- intervals[i].length == 2
 * 		- 0 <= li < ri <= 10^5
 * 		- All the given intervals are unique
 *
 * Examples:
 * 		Example 1:
 * 		Input: intervals = [[1,4],[3,6],[2,8]]
 * 		Output: 2
 * 		Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
 *
 * 		Example 2:
 * 		Input: intervals = [[1,4],[2,3]]
 * 		Output: 1
 */

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	var count = 0
	var prevEnd = 0

	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] > prevEnd {
			count++
			prevEnd = intervals[i][1]

		}
	}

	return count
}

func RunTestRemoveCoveredIntervals() {
	runner.InitMetrics("RemoveCoveredIntervals")

	testCases := map[string]struct {
		intervals [][]int
		expect    int
	}{
		"example-1": {
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			expect:    2,
		},
		"example-2": {
			intervals: [][]int{{1, 4}, {2, 3}},
			expect:    1,
		},
		"no-covered-intervals": {
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expect:    3,
		},
		"one-uncovered-interval": {
			intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			expect:    1,
		},
		"single-interval": {
			intervals: [][]int{{1, 4}},
			expect:    1,
		},
		"two-non-overlapping-intervals": {
			intervals: [][]int{{1, 3}, {4, 6}},
			expect:    2,
		},
		"same-start-different-end": {
			intervals: [][]int{{1, 4}, {1, 5}, {1, 3}},
			expect:    1,
		},
		"same-end-different-start": {
			intervals: [][]int{{1, 4}, {2, 4}, {3, 4}},
			expect:    1,
		},
		"completely-nested": {
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			expect:    1,
		},
		"adjacent-intervals": {
			intervals: [][]int{{1, 3}, {3, 5}, {5, 7}},
			expect:    3,
		},
		"multiple-independent-groups": {
			intervals: [][]int{{1, 5}, {2, 4}, {6, 10}, {7, 9}, {11, 15}},
			expect:    3,
		},
		"minimum-values": {
			intervals: [][]int{{0, 1}, {0, 2}},
			expect:    1,
		},
		"maximum-values": {
			intervals: [][]int{{0, 100000}, {10000, 50000}, {20000, 80000}},
			expect:    1,
		},
		"gap-intervals": {
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			expect:    4,
		},
		"partial-overlap": {
			intervals: [][]int{{1, 5}, {4, 8}},
			expect:    2,
		},
		"boundary-values": {
			intervals: [][]int{{0, 100000}, {1, 99999}},
			expect:    1,
		},
		"chain-coverage": {
			intervals: [][]int{{1, 6}, {2, 5}, {3, 4}},
			expect:    1,
		},
		"all-unique-non-overlapping": {
			intervals: [][]int{{1, 3}, {4, 6}, {7, 9}, {10, 12}},
			expect:    4,
		},
		"complex-case": {
			intervals: [][]int{{1, 4}, {0, 6}, {2, 3}, {5, 8}, {5, 7}},
			expect:    2,
		},
	}

	var passedCount int

	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"intervals": tc.intervals})

		result := runner.ExecCountMetrics(removeCoveredIntervals, tc.intervals).(int)
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
