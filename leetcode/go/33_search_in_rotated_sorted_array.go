package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Search in Rotated Sorted Array
Topic            : Array, Binary Search
Level            : Medium
URL              : https://leetcode.com/problems/search-in-rotated-sorted-array/
Description      : There is an integer array nums sorted in ascending order (with distinct values). Prior to being
        passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that
        the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
        For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2]. Given the array
        nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if
        it is not in nums.
Examples         :
        Example 1:
        Input: nums = [4,5,6,7,0,1,2], target = 0
        Output: 4

        Example 2:
        Input: nums = [4,5,6,7,0,1,2], target = 3
        Output: -1

        Example 3:
        Input: nums = [1], target = 0
        Output: -1
*/

func searchInRotatedSortedArray(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func RunTestSearchInRotatedSortedArray() {
	testCases := map[string]struct {
		nums   []int
		target int
		expect int
	}{
		"case-1": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		},
		"case-2": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			expect: -1,
		},
		"case-3": {
			nums:   []int{1},
			target: 0,
			expect: -1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := searchInRotatedSortedArray(testCase.nums, testCase.target)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
