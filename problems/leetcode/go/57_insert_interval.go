package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("InsertInterval", RunTestInsertInterval)
}

/**
 * Problem          : Insert Interval
 * Topics           : Array
 * Level            : Medium
 * URL              : https://leetcode.com/problems/insert-interval
 * Description      : You are given an array of non-overlapping intervals where intervals[i] = [starti, endi] represent
 *                    the start and the end of the ith interval and intervals is sorted in ascending order by starti.
 *                    You are also given an interval newInterval = [start, end] that represents the start and end of
 *                    another interval. Insert newInterval into intervals such that intervals is still sorted in
 *                    ascending order by starti and intervals still does not have any overlapping intervals (merge
 *                    overlapping intervals if necessary). Return intervals after the insertion.
 * Constraints      :
 *                    - 0 <= intervals.length <= 10^4
 *                    - intervals[i].length == 2
 *                    - 0 <= starti <= endi <= 10^5
 *                    - newInterval.length == 2
 *                    - 0 <= start <= end <= 10^5
 * Examples         :
 *                    Example 1:
 *                    Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 *                    Output: [[1,5],[6,9]]
 *
 *                    Example 2:
 *                    Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 *                    Output: [[1,2],[3,10],[12,16]]
 *                    Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
 */

func insertInterval_SolutionV2(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	newIntervals := make([][]int, 0, len(intervals)+1)
	i := 0

	for i < n && intervals[i][1] < newInterval[0] {
		newIntervals = append(newIntervals, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	newIntervals = append(newIntervals, newInterval)
	newIntervals = append(newIntervals, intervals[i:]...)

	return newIntervals
}

func insertInterval_SolutionV1(intervals [][]int, newInterval []int) [][]int {
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
	runner.InitMetrics("InsertInterval")

	testCases := map[string]struct {
		intervals   [][]int
		newInterval []int
		expect      [][]int
	}{
		"example-1-overlap-single": {
			intervals: [][]int{
				{1, 3}, {6, 9},
			},
			newInterval: []int{2, 5},
			expect: [][]int{
				{1, 5}, {6, 9},
			},
		},
		"example-2-overlap-multiple": {
			intervals: [][]int{
				{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
			},
			newInterval: []int{4, 8},
			expect: [][]int{
				{1, 2}, {3, 10}, {12, 16},
			},
		},
		"empty-intervals": {
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expect: [][]int{
				{5, 7},
			},
		},
		"insert-at-beginning-no-overlap": {
			intervals: [][]int{
				{3, 5}, {6, 7}, {8, 10},
			},
			newInterval: []int{1, 2},
			expect: [][]int{
				{1, 2}, {3, 5}, {6, 7}, {8, 10},
			},
		},
		"insert-at-end-no-overlap": {
			intervals: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
			newInterval: []int{7, 8},
			expect: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
		},
		"insert-in-middle-no-overlap": {
			intervals: [][]int{
				{1, 2}, {5, 6},
			},
			newInterval: []int{3, 4},
			expect: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
		"single-interval-overlap": {
			intervals: [][]int{
				{1, 5},
			},
			newInterval: []int{2, 3},
			expect: [][]int{
				{1, 5},
			},
		},
		"single-interval-no-overlap-before": {
			intervals: [][]int{
				{5, 7},
			},
			newInterval: []int{1, 3},
			expect: [][]int{
				{1, 3}, {5, 7},
			},
		},
		"single-interval-no-overlap-after": {
			intervals: [][]int{
				{1, 3},
			},
			newInterval: []int{5, 7},
			expect: [][]int{
				{1, 3}, {5, 7},
			},
		},
		"new-interval-covers-all": {
			intervals: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
			newInterval: []int{0, 10},
			expect: [][]int{
				{0, 10},
			},
		},
		"new-interval-inside-existing": {
			intervals: [][]int{
				{1, 10},
			},
			newInterval: []int{3, 5},
			expect: [][]int{
				{1, 10},
			},
		},
		"touching-intervals-start": {
			intervals: [][]int{
				{3, 5}, {6, 8},
			},
			newInterval: []int{1, 3},
			expect: [][]int{
				{1, 5}, {6, 8},
			},
		},
		"touching-intervals-end": {
			intervals: [][]int{
				{1, 2}, {3, 5},
			},
			newInterval: []int{5, 7},
			expect: [][]int{
				{1, 2}, {3, 7},
			},
		},
		"same-as-existing": {
			intervals: [][]int{
				{1, 3}, {4, 6},
			},
			newInterval: []int{1, 3},
			expect: [][]int{
				{1, 3}, {4, 6},
			},
		},
		"merge-all-consecutive": {
			intervals: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
			newInterval: []int{2, 7},
			expect: [][]int{
				{1, 8},
			},
		},
		"new-interval-merges-first-two": {
			intervals: [][]int{
				{1, 2}, {4, 5}, {7, 8},
			},
			newInterval: []int{2, 4},
			expect: [][]int{
				{1, 5}, {7, 8},
			},
		},
		"new-interval-merges-last-two": {
			intervals: [][]int{
				{1, 2}, {3, 4}, {6, 7},
			},
			newInterval: []int{4, 6},
			expect: [][]int{
				{1, 2}, {3, 7},
			},
		},
		"point-interval": {
			intervals: [][]int{
				{1, 5},
			},
			newInterval: []int{5, 5},
			expect: [][]int{
				{1, 5},
			},
		},
		"large-numbers": {
			intervals: [][]int{
				{10000, 10000}, {99999, 99999},
			},
			newInterval: []int{50000, 60000},
			expect: [][]int{
				{10000, 10000}, {50000, 60000}, {99999, 99999},
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := runner.ExecCountMetrics(insertInterval_SolutionV2, testCase.intervals, testCase.newInterval).([][]int)
		format.PrintInput(map[string]interface{}{"intervals": testCase.intervals, "newInterval": testCase.newInterval})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
