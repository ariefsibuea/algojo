package main

import (
	"fmt"
	"os"
)

/**
 * LeetCode Problem : Trapping Rain Water
 * Topics           : Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack
 * Level            : Hard
 * URL              : https://leetcode.com/problems/trapping-rain-water
 * Description      : Given n non-negative integers representing an elevation map where the width of each bar is 1,
 * 					compute how much water it can trap after raining.
 * Examples         :
 * 					Example 1:
 * 					Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 					Output: 6
 * 					Explanation: The above elevation map (black section) is represented by array
 * 					[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
 *
 * 					Example 2:
 * 					Input: height = [4,2,0,3,2,5]
 * 					Output: 9
 * Reference		: https://www.hellointerview.com/learn/code/two-pointers/trapping-rain-water
 */

func trappingRainWater(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	count := 0

	for left < right {
		if maxLeft < maxRight {
			left += 1
			if maxLeft < height[left] {
				maxLeft = height[left]
			} else {
				count += maxLeft - height[left]
			}
		} else {
			right -= 1
			if maxRight < height[right] {
				maxRight = height[right]
			} else {
				count += maxRight - height[right]
			}
		}
	}

	return count
}

func RunTestTrappingRainWater() {
	testCases := map[string]struct {
		height []int
		expect int
	}{
		"case-1": {
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expect: 6,
		},
		"case-2": {
			height: []int{4, 2, 0, 3, 2, 5},
			expect: 9,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := trappingRainWater(testCase.height)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}

	fmt.Printf("\n✅ All tests passed!\n")
}
