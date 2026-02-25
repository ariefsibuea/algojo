package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("MeetingRooms", RunTestMeetingRooms)
}

/**
 * Problem 			: Meeting Rooms
 * Topics           : Array, Sorting
 * Level            : Easy
 * URL              : https://neetcode.io/problems/meeting-schedule?list=neetcode150
 * Description      : Given an array of meeting time interval objects consisting of Start and end times
 * 					[[Start_1,end_1],[Start_2,end_2],...] (Start_i < end_i), determine if a person could add all
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

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var previous = intervals[0]
	var current Interval

	for i := 1; i < len(intervals); i++ {
		current = intervals[i]
		if previous.End > current.Start {
			return false
		}
		previous = current
	}

	return true
}

func RunTestMeetingRooms() {
	testCases := map[string]struct {
		intervals []Interval
		expect    bool
	}{
		"case-1": {
			intervals: []Interval{
				{Start: 0, End: 30},
				{Start: 5, End: 10},
				{Start: 15, End: 20},
			},
			expect: false,
		},
		"case-2": {
			intervals: []Interval{
				{Start: 5, End: 8},
				{Start: 9, End: 15},
			},
			expect: true,
		},
		"case-3": {
			intervals: []Interval{
				{Start: 5, End: 8},
				{Start: 8, End: 15},
			},
			expect: true,
		},
		"case-4": {
			intervals: []Interval{
				{Start: 5, End: 8},
				{Start: 7, End: 15},
			},
			expect: false,
		},
		"case-5": {
			intervals: []Interval{
				{Start: 0, End: 15},
				{Start: 15, End: 30},
				{Start: 30, End: 45},
				{Start: 45, End: 60},
				{Start: 60, End: 75},
				{Start: 75, End: 90},
				{Start: 85, End: 100},
			},
			expect: false,
		},
		"case-6": {
			intervals: []Interval{
				{Start: 0, End: 15},
				{Start: 0, End: 1},
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
