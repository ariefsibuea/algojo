package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/**
 * Problem          : Merge Intervals
 * Topics           : Array, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/merge-intervals
 * Description      : You are given an array of intervals where each interval is represented as [start_i, end_i].
 *                    Your task is to merge all overlapping intervals and return an array containing only
 * 					  non-overlapping intervals that cover all the original intervals. Two intervals overlap if one
 * 					  starts before or when the other ends. For example, [1,3] and [2,6] overlap and should be merged
 * 					  into [1,6], while [1,2] and [3,4] do not overlap and remain separate.
 * Constraints      :
 *                    - 1 <= intervals.length <= 10^4
 *                    - intervals[i].length == 2
 *                    - 0 <= starti <= endi <= 10^4
 * Examples         :
 *                    Example 1:
 *                    Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 *                    Output: [[1,6],[8,10],[15,18]]
 *                    Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
 *
 *                    Example 2:
 *                    Input: intervals = [[1,4],[4,5]]
 *                    Output: [[1,5]]
 *                    Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 *                    Example 3:
 *                    Input: intervals = [[4,7],[1,4]]
 *                    Output: [[1,7]]
 *                    Explanation: Intervals [1,4] and [4,7] are considered overlapping.
 */

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] <= end {
			end = max(end, interval[1])
		} else {
			result = append(result, []int{start, end})
			start, end = interval[0], interval[1]
		}
	}
	result = append(result, []int{start, end})

	return result
}

func RunTestMergeIntervals() {
	runner.InitMetrics("MergeIntervals")

	testCases := map[string]struct {
		intervals [][]int
		expect    [][]int
	}{
		"example-1-basic-overlap": {
			intervals: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			expect: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
		"example-2-touching-intervals": {
			intervals: [][]int{
				{1, 4}, {4, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
		"example-3-unsorted-input": {
			intervals: [][]int{
				{4, 7}, {1, 4},
			},
			expect: [][]int{
				{1, 7},
			},
		},
		"single-interval": {
			intervals: [][]int{
				{1, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
		"no-overlap-multiple-intervals": {
			intervals: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
			expect: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
		},
		"all-overlapping-into-one": {
			intervals: [][]int{
				{1, 4}, {2, 5}, {3, 6}, {4, 7},
			},
			expect: [][]int{
				{1, 7},
			},
		},
		"nested-intervals": {
			intervals: [][]int{
				{1, 10}, {2, 3}, {4, 5}, {6, 7},
			},
			expect: [][]int{
				{1, 10},
			},
		},
		"same-start-different-end": {
			intervals: [][]int{
				{1, 3}, {1, 5}, {1, 2},
			},
			expect: [][]int{
				{1, 5},
			},
		},
		"same-intervals": {
			intervals: [][]int{
				{1, 4}, {1, 4}, {1, 4},
			},
			expect: [][]int{
				{1, 4},
			},
		},
		"large-numbers": {
			intervals: [][]int{
				{10000, 10000}, {9999, 10000}, {0, 1},
			},
			expect: [][]int{
				{0, 1}, {9999, 10000},
			},
		},
		"empty-intervals-array": {
			intervals: [][]int{},
			expect:    [][]int{},
		},
		"point-intervals": {
			intervals: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
			expect: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
		},
		"chain-overlap": {
			intervals: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {4, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		// Make a copy of intervals to preserve original for display
		intervalsCopy := make([][]int, len(testCase.intervals))
		copy(intervalsCopy, testCase.intervals)

		result := runner.ExecCountMetrics(mergeIntervals, testCase.intervals).([][]int)
		format.PrintInput(map[string]interface{}{"intervals": testCase.intervals})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
