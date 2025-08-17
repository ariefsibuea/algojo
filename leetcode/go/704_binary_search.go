package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Binary Search
Topic            : Array, Binary Search
Level            : Easy
URL              : https://leetcode.com/problems/binary-search/
Description      : Given an array of integers nums which is sorted in ascending order, and an integer target, write a
        function to search target in nums. If target exists, then return its index. Otherwise, return -1. You must
        write an algorithm with O(log n) runtime complexity.
Examples         :
        Example 1:
        Input: nums = [-1,0,3,5,9,12], target = 9
        Output: 4
        Explanation: 9 exists in nums and its index is 4

        Example 2:
        Input: nums = [-1,0,3,5,9,12], target = 2
        Output: -1
        Explanation: 2 does not exist in nums so return -1
*/

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			start = mid + 1
		default:
			end = mid - 1
		}
	}

	return -1
}

func RunTestBinarySearch() {
	testCases := map[string]struct {
		nums   []int
		target int
		expect int
	}{
		"case-1": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			expect: 4,
		},
		"case-2": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			expect: -1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := binarySearch(testCase.nums, testCase.target)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
