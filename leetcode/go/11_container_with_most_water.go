package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Container With Most Water
Topic            : Array, Two Pointers, Greedy
Level            : Medium
URL              : https://leetcode.com/problems/container-with-most-water
Description      : You are given an integer array height of length n. There are n vertical lines drawn such that the two
        endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a
        container, such that the container contains the most water. Return the maximum amount of water a container can
        store. Notice that you may not slant the container.
Examples         :
        Example 1:
        Input: height = [1,8,6,2,5,4,8,3,7]
        Output: 49
        Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area
                of water (blue section) the container can contain is 49.

        Example 2:
        Input: height = [1,1]
        Output: 1
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

func RunTestMaxArea() {
	testCases := map[string]struct {
		height []int
		expect int
	}{
		"case-1": {
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expect: 49,
		},
		"case-2": {
			height: []int{1, 1},
			expect: 1,
		},
	}

	var result int

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result = maxArea(testCase.height)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
