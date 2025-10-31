package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Course Schedule
 * Topic            : Depth-First Search, Breadth-First Search, Graph, Topological Sort
 * Level            : Medium
 * URL              : https://leetcode.com/problems/course-schedule
 */

func canFinishCourseSchedule(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		dest, src := prerequisite[0], prerequisite[1]
		graph[src] = append(graph[src], dest)
		inDegree[dest]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func RunTestCourseSchedule() {
	testCases := map[string]struct {
		numCourses    int
		prerequisites [][]int
		expect        bool
	}{
		"case-1": {
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expect:        true,
		},
		"case-2": {
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expect:        false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := canFinishCourseSchedule(testCase.numCourses, testCase.prerequisites)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}
	fmt.Printf("\n✅ All tests passed!\n")
}
