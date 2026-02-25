package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SearchInRotatedSortedArray", RunTestSearchInRotatedSortedArray)
}

/*
 * Problem	: Search in Rotated Sorted Array
 * Topics	: Array, Binary Search
 * Level	: Medium
 * URL		: https://leetcode.com/problems/search-in-rotated-sorted-array/
 *
 * Description:
 * 		There is an integer array nums sorted in ascending order (with distinct values). Prior to being
 * 		passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.lengt)
 * 		such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ...,
 * 		nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become
 * 		[4,5,6,7,0,1,2]. Given the array nums after the possible rotation and an integer target, return
 * 		the index of target if it is in nums, or -1 if it is not in nums.
 *
 * Constraints:
 * 		- 1 <= nums.length <= 5000
 * 		- -10^4 <= nums[i] <= 10^4
 * 		- All values of nums are unique.
 * 		- nums is an ascending array that is possibly rotated.
 * 		- -10^4 <= target <= 10^4
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [4,5,6,7,0,1,2], target = 0
 * 		Output: 4
 *
 * 		Example 2:
 * 		Input: nums = [4,5,6,7,0,1,2], target = 3
 * 		Output: -1
 *
 * 		Example 3:
 * 		Input: nums = [1], target = 0
 * 		Output: -1
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
	runner.InitMetrics("SearchInRotatedSortedArray")

	testCases := map[string]struct {
		nums   []int
		target int
		expect int
	}{
		"example-1-found": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		},
		"example-2-not-found": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			expect: -1,
		},
		"example-3-single-element": {
			nums:   []int{1},
			target: 0,
			expect: -1,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums, "target": tc.target})

		result := runner.ExecCountMetrics(searchInRotatedSortedArray, tc.nums, tc.target).(int)
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
