package main

import (
	"fmt"
	"os"
	"sort"
)

/**
 * Problem 			: Merge Intervals
 * Topics           : Array, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/merge-intervals
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 					Output: [[1,6],[8,10],[15,18]]
 * 					Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
 *
 * 					Example 2:
 * 					Input: intervals = [[1,4],[4,5]]
 * 					Output: [[1,5]]
 * 					Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 * 					Example 3:
 * 					Input: intervals = [[4,7],[1,4]]
 * 					Output: [[1,7]]
 * 					Explanation: Intervals [1,4] and [4,7] are considered overlapping.
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
	testCases := map[string]struct {
		intervals [][]int
		expect    [][]int
	}{
		"case-1": {
			intervals: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			expect: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
		"case-2": {
			intervals: [][]int{
				{1, 4}, {4, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
		"case-3": {
			intervals: [][]int{
				{4, 7}, {1, 4},
			},
			expect: [][]int{
				{1, 7},
			},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := mergeIntervals(testCase.intervals)
		if !EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
