package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Meeting Rooms
 * Topics           : Array, Sorting
 * Level            : Easy
 * URL              : https://neetcode.io/problems/meeting-schedule?list=neetcode150
 * Description      : Given an array of meeting time interval objects consisting of start and end times
 * 					[[start_1,end_1],[start_2,end_2],...] (start_i < end_i), determine if a person could add all
 * 					meetings to their schedule without any conflicts.
 * 					Note: (0,8),(8,10) is not considered a conflict at 8
 * Examples         :
 * 					Example 1:
 * 					Input: intervals = [(0,30),(5,10),(15,20)]
 * 					Output: false
 *
 * 					Example 2:
 * 					Input: intervals = [(5,8),(9,15)]
 * 					Output: true
 */

type IntervalP252 struct {
	start int
	end   int
}

func canAttendMeetings(intervals []IntervalP252) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var previous = intervals[0]
	var current IntervalP252

	for i := 1; i < len(intervals); i++ {
		current = intervals[i]
		if previous.end > current.start {
			return false
		}
		previous = current
	}

	return true
}

func RunTestMeetingRooms() {
	testCases := map[string]struct {
		intervals []IntervalP252
		expect    bool
	}{
		"case-1": {
			intervals: []IntervalP252{
				{start: 0, end: 30},
				{start: 5, end: 10},
				{start: 15, end: 20},
			},
			expect: false,
		},
		"case-2": {
			intervals: []IntervalP252{
				{start: 5, end: 8},
				{start: 9, end: 15},
			},
			expect: true,
		},
		"case-3": {
			intervals: []IntervalP252{
				{start: 5, end: 8},
				{start: 8, end: 15},
			},
			expect: true,
		},
		"case-4": {
			intervals: []IntervalP252{
				{start: 5, end: 8},
				{start: 7, end: 15},
			},
			expect: false,
		},
		"case-5": {
			intervals: []IntervalP252{
				{start: 0, end: 15},
				{start: 15, end: 30},
				{start: 30, end: 45},
				{start: 45, end: 60},
				{start: 60, end: 75},
				{start: 75, end: 90},
				{start: 85, end: 100},
			},
			expect: false,
		},
		"case-6": {
			intervals: []IntervalP252{
				{start: 0, end: 15},
				{start: 0, end: 1},
			},
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := canAttendMeetings(testCase.intervals)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
