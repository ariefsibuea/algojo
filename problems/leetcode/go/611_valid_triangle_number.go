package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("ValidTriangleNumber", RunTestValidTriangleNumber)
}

/**
 * LeetCode Problem : Valid Triangle Number
 * Topic            : Array, Two Pointers, Binary Search, Greedy, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/valid-triangle-number
 * Description      : Given an integer array nums, return the number of triplets chosen from the array that can make
 * 					triangles if we take them as side lengths of a triangle.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [2,2,3,4]
 * 					Output: 3
 * 					Explanation: Valid combinations are:
 * 					2,3,4 (using the first 2)
 * 					2,3,4 (using the second 2)
 * 					2,2,3
 *
 * 					Example 2:
 * 					Input: nums = [4,2,3,4]
 * 					Output: 4
 */

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	count := 0

	for i := len - 1; i > 1; i-- {
		n := nums[i]
		left := 0
		right := i - 1

		for left < right {
			// NOTE: A triangle is valid if a + b > c, b + c > a, a + c > b.
			if nums[left]+nums[right] > n {
				count += right - left
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return count
}

func RunTestValidTriangleNumber() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{2, 2, 3, 4},
			expect: 3,
		},
		"case-2": {
			nums:   []int{4, 2, 3, 4},
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := triangleNumber(testCase.nums)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
