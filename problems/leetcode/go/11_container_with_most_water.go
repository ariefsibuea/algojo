package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: Container With Most Water
 * Topics	: Array, Two Pointers, Greedy
 * Level	: Medium
 * URL		: https://leetcode.com/problems/container-with-most-water/
 *
 * Description:
 * 		You are given an integer array height of length n. There are n vertical lines drawn such that the
 * 		two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the
 * 		x-axis form a container, such that the container contains the most water. Return the maximum
 * 		amount of water a container can store. Notice that you may not slant the container.
 *
 * Constraints:
 * 		- n == height.length
 * 		- 2 <= n <= 10^5
 * 		- 0 <= height[i] <= 10^4
 *
 * Examples:
 * 		Example 1:
 * 		Input: height = [1,8,6,2,5,4,8,3,7]
 * 		Output: 49
 * 		Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
 * 		the max area of water (blue section) the container can contain is 49.
 *
 * 		Example 2:
 * 		Input: height = [1,1]
 * 		Output: 1
 */

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	maxAmount := -1

	for start < end {
		minHeight := min(height[start], height[end])
		maxAmount = max(minHeight*(end-start), maxAmount)

		if height[start] < height[end] {
			start = start + 1
		} else {
			end = end - 1
		}
	}

	return maxAmount
}

func RunTestContainerWithMostWater() {
	runner.InitMetrics("ContainerWithMostWater")

	testCases := map[string]struct {
		height []int
		expect int
	}{
		"example-1-basic": {
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expect: 49,
		},
		"example-2-minimal": {
			height: []int{1, 1},
			expect: 1,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"height": tc.height})

		result := runner.ExecCountMetrics(maxArea, tc.height).(int)
		if !cmp.EqualNumbers(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
