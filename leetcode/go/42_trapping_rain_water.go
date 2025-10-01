package main

/**
 * LeetCode Problem : Trapping Rain Water
 * Topics           : Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack
 * Level            : Hard
 * URL              : https://leetcode.com/problems/trapping-rain-water
 * Description      :
 * Examples         :
 * 			Example 1:
 * 			Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 			Output: 6
 * 			Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In
 * 			this case, 6 units of rain water (blue section) are being trapped.
 *
 * 			Example 2:
 * 			Input: height = [4,2,0,3,2,5]
 * 			Output: 9
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
