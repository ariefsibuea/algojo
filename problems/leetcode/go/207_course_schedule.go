package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Course Schedule
 * Topics           : Depth-First Search, Breadth-First Search, Graph, Topological Sort
 * Level            : Medium
 * URL              : https://leetcode.com/problems/course-schedule
 * Description      : There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
 *                    You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must
 * 					  take course bi first if you want to take course ai.
 *                    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
 *                    Return true if you can finish all courses. Otherwise, return false.
 * Constraints      :
 *                    - 1 <= numCourses <= 2000
 *                    - 0 <= prerequisites.length <= 5000
 *                    - prerequisites[i].length == 2
 *                    - 0 <= ai, bi < numCourses
 *                    - All the pairs prerequisites[i] are unique.
 * Examples         :
 *                    Example 1:
 *                    Input: numCourses = 2, prerequisites = [[1,0]]
 *                    Output: true
 *                    Explanation: There are a total of 2 courses to take.
 *                    To take course 1 you should have finished course 0. So it is possible.
 *
 *                    Example 2:
 *                    Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 *                    Output: false
 *                    Explanation: There are a total of 2 courses to take.
 *                    To take course 1 you should have finished course 0, and to take course 0 you should also have
 * 					  finished course 1. So it is impossible.
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
