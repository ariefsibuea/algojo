package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: Count Days Without Meetings
 * Topics	: Array, Sorting, Interval Merging
 * Level	: Medium
 * URL		: https://leetcode.com/problems/count-days-without-meetings
 *
 * Description:
 * 		You are given a positive integer 'days' representing the total number of days an employee is available for work
 * 		(days are numbered from 1 to days). You are also given a 2D array 'meetings' where each element
 * 		meetings[i] = [start_i, end_i] represents a meeting that runs from start_i to end_i (both days inclusive). Your
 * 		task is to count the total number of days when the employee is available for work but no meetings are
 * 		scheduled. Note that meetings may overlap with each other.
 *
 * Constraints:
 * 		- 1 <= days <= 10^9
 * 		- 1 <= meetings.length <= 10^5
 * 		- meetings[i].length == 2
 * 		- 1 <= meetings[i][0] <= meetings[i][1] <= days
 *
 * Examples:
 * 		Example 1:
 * 		Input: days = 10, meetings = [[5,7],[1,3],[9,10]]
 * 		Output: 2
 * 		Explanation: There is no meeting scheduled on the 4th and 8th days. Days 1-3 have meetings (from
 * 		meeting [1,3]), days 5-7 have meetings (from meeting [5,7]), and days 9-10 have meetings (from
 * 		meeting [9,10]). The free days are day 4 and day 8, totaling 2 days.
 *
 * 		Example 2:
 * 		Input: days = 5, meetings = [[2,4],[1,3]]
 * 		Output: 1
 * 		Explanation: Meetings [1,3] and [2,4] overlap, covering days 1-4. Only day 5 is free.
 *
 * 		Example 3:
 * 		Input: days = 6, meetings = [[1,6]]
 * 		Output: 0
 * 		Explanation: The single meeting covers all days from 1 to 6, so there are no free days.
 */

func countDays(days int, meetings [][]int) int {
	if len(meetings) == 0 {
		return days
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] != meetings[j][0] {
			return meetings[i][0] < meetings[j][0]
		}
		return meetings[i][1] < meetings[j][1]
	})

	var start, end = meetings[0][0], meetings[0][1]
	var meetingDays = 0

	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= end {
			end = max(end, meetings[i][1])
			continue
		}

		meetingDays += end - start + 1
		start = meetings[i][0]
		end = meetings[i][1]
	}

	meetingDays += end - start + 1

	return days - meetingDays
}

func RunTestCountDaysWithoutMeetings() {
	runner.InitMetrics("CountDaysWithoutMeetings")

	testCases := map[string]struct {
		days     int
		meetings [][]int
		expect   int
	}{
		"example-1-multiple-meetings-with-gaps": {
			days: 10,
			meetings: [][]int{
				{5, 7}, {1, 3}, {9, 10},
			},
			expect: 2,
		},
		"example-2-overlapping-meetings": {
			days: 5,
			meetings: [][]int{
				{2, 4}, {1, 3},
			},
			expect: 1,
		},
		"example-3-full-coverage": {
			days: 6,
			meetings: [][]int{
				{1, 6},
			},
			expect: 0,
		},
		"edge-case-no-meetings": {
			days:     10,
			meetings: [][]int{},
			expect:   10,
		},
		"edge-case-single-day-no-meeting": {
			days:     1,
			meetings: [][]int{},
			expect:   1,
		},
		"edge-case-meeting-at-end-only": {
			days: 5,
			meetings: [][]int{
				{5, 5},
			},
			expect: 4,
		},
		"edge-case-consecutive-meetings": {
			days: 5,
			meetings: [][]int{
				{1, 2}, {3, 4},
			},
			expect: 1,
		},
		"multiple-overlapping-meetings": {
			days: 8,
			meetings: [][]int{
				{1, 3}, {2, 5}, {4, 7},
			},
			expect: 1,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"days": testCase.days, "meetings": testCase.meetings})

		result := runner.ExecCountMetrics(countDays, testCase.days, testCase.meetings).(int)
		if !cmp.IsEqual(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
