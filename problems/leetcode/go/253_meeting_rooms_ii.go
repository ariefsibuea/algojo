package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem          : Meeting Rooms II
 * Topics           : Greedy, Array, Two Pointers, Prefix Sum, Sorting, Heap (Priority Queue)
 * Level            : Medium
 * URL              : https://leetcode.com/problems/meeting-rooms-ii/
 * Description      : Given an array of meeting time intervals where intervals[i] = [start_i, end_i],
 *                    find the minimum number of conference rooms required to schedule all meetings
 *                    without conflicts. When meetings overlap in time, they require separate rooms.
 * Constraints      :
 *                    - 1 <= intervals.length <= 10^4
 *                    - 0 <= start_i < end_i <= 10^6
 * Examples         :
 *                    Example 1:
 *                    Input: intervals = [[0,30],[5,10],[15,20]]
 *                    Output: 2
 *                    Explanation:
 *                    - Meeting 1 (0-30) requires room 1
 *                    - Meeting 2 (5-10) overlaps with meeting 1, requires room 2
 *                    - Meeting 3 (15-20) overlaps with meeting 1, but meeting 2 has ended, so reuse room 2
 *                    Minimum 2 rooms needed.
 *
 *                    Example 2:
 *                    Input: intervals = [[7,10],[2,4]]
 *                    Output: 1
 *                    Explanation:
 *                    - Meeting 1 (7-10) requires room 1
 *                    - Meeting 2 (2-4) ends before meeting 1 starts, same room can be reused
 *                    Minimum 1 room needed.
 */

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)
	actions := make([][2]int, 0, 2*len(intervals))

	for i := 0; i < n; i++ {
		actions = append(actions, [2]int{intervals[i].Start, 1}, [2]int{intervals[i].End, -1})
	}

	sort.Slice(actions, func(i, j int) bool {
		if actions[i][0] != actions[j][0] {
			return actions[i][0] < actions[j][0]
		}
		return actions[i][1] < actions[j][1]
	})

	maxRooms, numRooms := 0, 0

	for i := 0; i < len(actions); i++ {
		numRooms += actions[i][1]
		maxRooms = max(maxRooms, numRooms)
	}

	return maxRooms
}

func RunTestMeetingRoomsII() {
	runner.InitMetrics("MeetingRoomsII")

	testCases := map[string]struct {
		intervals []Interval
		expect    int
	}{
		"example-1-overlapping-meetings": {
			intervals: []Interval{
				{Start: 0, End: 30},
				{Start: 5, End: 10},
				{Start: 15, End: 20},
			},
			expect: 2,
		},
		"example-2-non-overlapping-meetings": {
			intervals: []Interval{
				{Start: 7, End: 10},
				{Start: 2, End: 4},
			},
			expect: 1,
		},
		"single-meeting": {
			intervals: []Interval{
				{Start: 0, End: 10},
			},
			expect: 1,
		},
		"all-overlapping-same-time": {
			intervals: []Interval{
				{Start: 0, End: 10},
				{Start: 0, End: 10},
				{Start: 0, End: 10},
				{Start: 0, End: 10},
			},
			expect: 4,
		},
		"sequential-no-overlap": {
			intervals: []Interval{
				{Start: 0, End: 5},
				{Start: 5, End: 10},
				{Start: 10, End: 15},
				{Start: 15, End: 20},
			},
			expect: 1,
		},
		"chain-overlap": {
			intervals: []Interval{
				{Start: 0, End: 30},
				{Start: 5, End: 35},
				{Start: 10, End: 40},
			},
			expect: 3,
		},
		"touching-boundaries": {
			intervals: []Interval{
				{Start: 0, End: 5},
				{Start: 5, End: 10},
			},
			expect: 1,
		},
		"nested-intervals": {
			intervals: []Interval{
				{Start: 0, End: 100},
				{Start: 10, End: 20},
				{Start: 30, End: 40},
			},
			expect: 2,
		},
		"multiple-rooms-required": {
			intervals: []Interval{
				{Start: 0, End: 10},
				{Start: 5, End: 15},
				{Start: 10, End: 20},
				{Start: 15, End: 25},
			},
			expect: 2,
		},
		"large-gap-between": {
			intervals: []Interval{
				{Start: 0, End: 5},
				{Start: 1000, End: 1005},
				{Start: 2000, End: 2005},
			},
			expect: 1,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := runner.ExecCountMetrics(minMeetingRooms, testCase.intervals).(int)
		format.PrintInput(map[string]interface{}{"intervals": testCase.intervals})

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
