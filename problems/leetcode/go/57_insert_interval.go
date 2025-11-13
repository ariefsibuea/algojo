package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Insert Interval
 * Topics           : Array
 * Level            : Medium
 * URL              : https://leetcode.com/problems/insert-interval
 * Description      : You are given an array of non-overlapping intervals intervals where
 * 					intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is
 * 					sorted in ascending order by starti. You are also given an interval newInterval = [start, end]
 * 					that represents the start and end of another interval.
 *
 * 					Insert newInterval into intervals such that intervals is still sorted in ascending order by starti
 * 					and intervals still does not have any overlapping intervals (merge overlapping intervals if
 * 					necessary). Return intervals after the insertion.
 * Examples         :
 * 					Example 1:
 * 					Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 					Output: [[1,5],[6,9]]
 *
 * 					Example 2:
 * 					Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 					Output: [[1,2],[3,10],[12,16]]
 * 					Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
 */

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var result = make([][]int, 0, len(intervals)+1)

	var interval []int
	var index int

intervalLoop:
	for index < len(intervals) {
		interval = intervals[index]

		switch {
		case interval[0] > newInterval[1]:
			result = append(result, newInterval)
			result = append(result, intervals[index:]...)
			newInterval = nil
			break intervalLoop
		case interval[0] <= newInterval[1] && newInterval[0] <= interval[1]:
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		default:
			result = append(result, interval)
		}

		index += 1
	}

	if len(newInterval) != 0 {
		result = append(result, newInterval)
	}

	return result
}

func RunTestInsertInterval() {
	testCases := map[string]struct {
		intervals   [][]int
		newInterval []int
		expect      [][]int
	}{
		"case-1": {
			intervals: [][]int{
				{1, 2}, {5, 6},
			},
			newInterval: []int{3, 4},
			expect: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
		"case-2": {
			intervals: [][]int{
				{1, 3}, {6, 9},
			},
			newInterval: []int{2, 5},
			expect: [][]int{
				{1, 5}, {6, 9},
			},
		},
		"case-3": {
			intervals: [][]int{
				{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
			},
			newInterval: []int{4, 8},
			expect: [][]int{
				{1, 2}, {3, 10}, {12, 16},
			},
		},
		"case-5": {
			intervals: [][]int{
				{1, 2}, {3, 4},
			},
			newInterval: []int{5, 6},
			expect: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := insertInterval(testCase.intervals, testCase.newInterval)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
