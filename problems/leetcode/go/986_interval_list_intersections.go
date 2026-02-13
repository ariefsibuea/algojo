package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem          : Interval List Intersections
 * Topics           : Array, Two Pointers
 * Level            : Medium
 * URL              : https://leetcode.com/problems/interval-list-intersections
 * Description      : You are given two lists of closed intervals, firstList and secondList, where each interval is
 *                    represented as [start, end]. Each list contains intervals that are pairwise disjoint (no two
 *                    intervals within the same list overlap) and sorted in ascending order by their start time. A
 *                    closed interval [a, b] denotes the set of real numbers x where a <= x <= b. The task is to find
 *                    and return the intersection of these two interval lists. An intersection between two intervals is
 *                    the set of numbers that appear in both intervals, which is either empty or can be represented as
 *                    another closed interval.
 * Constraints      :
 *                    - 0 <= firstList.length, secondList.length <= 1000
 *                    - firstList.length + secondList.length >= 1
 *                    - 0 <= starti < endi <= 10^9
 *                    - endi < starti+1 (intervals in firstList are disjoint)
 *                    - 0 <= startj < endj <= 10^9
 *                    - endj < startj+1 (intervals in secondList are disjoint)
 * Examples         :
 *                    Example 1:
 *                    Input: firstList = [[0,2],[5,10],[13,23],[24,25]],
 *                           secondList = [[1,5],[8,12],[15,24],[25,26]]
 *                    Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 *
 *                    Example 2:
 *                    Input: firstList = [[1,3],[5,9]], secondList = []
 *                    Output: []
 */

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intersectList := make([][]int, 0, len(firstList))

	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		if firstList[i][0] <= secondList[j][1] && firstList[i][1] >= secondList[j][0] {
			intersect := make([]int, 2)
			intersect[0] = max(firstList[i][0], secondList[j][0])
			intersect[1] = min(firstList[i][1], secondList[j][1])
			intersectList = append(intersectList, intersect)
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return intersectList
}

func intervalIntersection_SolutionV1(firstList [][]int, secondList [][]int) [][]int {
	actions := make([][2]int, 0, 2*(len(firstList)+len(secondList)))

	for i := 0; i < len(firstList); i++ {
		actions = append(actions, [2]int{firstList[i][0], 1}, [2]int{firstList[i][1], -1})
	}

	for i := 0; i < len(secondList); i++ {
		actions = append(actions, [2]int{secondList[i][0], 1}, [2]int{secondList[i][1], -1})
	}

	sort.Slice(actions, func(i, j int) bool {
		if actions[i][0] != actions[j][0] {
			return actions[i][0] < actions[j][0]
		}
		return actions[i][1] > actions[j][1]
	})

	var n = 0
	var result = make([][]int, 0)
	var intersect []int

	for i := 0; i < len(actions); i++ {
		n += actions[i][1]
		if n == 2 && actions[i][1] == 1 {
			intersect = make([]int, 2)
			intersect[0] = actions[i][0]
		}
		if n == 1 && actions[i][1] == -1 {
			intersect[1] = actions[i][0]
			result = append(result, intersect)
		}
	}

	return result
}

func RunTestIntervalListIntersections() {
	runner.InitMetrics("IntervalListIntersections")

	testCases := map[string]struct {
		firstList  [][]int
		secondList [][]int
		expect     [][]int
	}{
		"example-1-multiple-overlaps": {
			firstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			expect:     [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		"example-2-empty-second-list": {
			firstList:  [][]int{{1, 3}, {5, 9}},
			secondList: [][]int{},
			expect:     [][]int{},
		},
		"empty-first-list": {
			firstList:  [][]int{},
			secondList: [][]int{{1, 3}, {5, 9}},
			expect:     [][]int{},
		},
		"both-empty-lists": {
			firstList:  [][]int{},
			secondList: [][]int{},
			expect:     [][]int{},
		},
		"no-intersections": {
			firstList:  [][]int{{1, 3}, {7, 9}},
			secondList: [][]int{{4, 6}, {10, 12}},
			expect:     [][]int{},
		},
		"single-interval-each-complete-overlap": {
			firstList:  [][]int{{1, 10}},
			secondList: [][]int{{3, 8}},
			expect:     [][]int{{3, 8}},
		},
		"single-interval-each-partial-overlap": {
			firstList:  [][]int{{1, 5}},
			secondList: [][]int{{3, 10}},
			expect:     [][]int{{3, 5}},
		},
		"touching-boundaries": {
			firstList:  [][]int{{1, 3}, {5, 7}},
			secondList: [][]int{{3, 5}, {7, 9}},
			expect:     [][]int{{3, 3}, {5, 5}, {7, 7}},
		},
		"complete-containment": {
			firstList:  [][]int{{1, 10}},
			secondList: [][]int{{2, 3}, {4, 6}, {7, 9}},
			expect:     [][]int{{2, 3}, {4, 6}, {7, 9}},
		},
		"first-list-contained-in-second": {
			firstList:  [][]int{{2, 3}, {5, 6}},
			secondList: [][]int{{1, 10}},
			expect:     [][]int{{2, 3}, {5, 6}},
		},
		"multiple-intervals-dense-overlap": {
			firstList:  [][]int{{1, 2}, {3, 4}, {5, 6}},
			secondList: [][]int{{1, 6}},
			expect:     [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		"large-values": {
			firstList:  [][]int{{0, 1000000000}},
			secondList: [][]int{{500000000, 600000000}},
			expect:     [][]int{{500000000, 600000000}},
		},
		"single-point-overlap": {
			firstList:  [][]int{{1, 2}, {3, 4}},
			secondList: [][]int{{2, 3}, {4, 5}},
			expect:     [][]int{{2, 2}, {3, 3}, {4, 4}},
		},
		"adjacent-intervals-no-overlap": {
			firstList:  [][]int{{1, 2}, {4, 5}},
			secondList: [][]int{{2, 4}},
			expect:     [][]int{{2, 2}, {4, 4}},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		firstCopy := make([][]int, len(testCase.firstList))
		copy(firstCopy, testCase.firstList)
		secondCopy := make([][]int, len(testCase.secondList))
		copy(secondCopy, testCase.secondList)

		result := runner.ExecCountMetrics(intervalIntersection, firstCopy, secondCopy).([][]int)
		format.PrintInput(map[string]interface{}{"firstList": testCase.firstList, "secondList": testCase.secondList})

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
