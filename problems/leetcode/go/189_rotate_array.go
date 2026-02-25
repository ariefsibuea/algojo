package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("RotateArray", RunTestRotateArray)
}

/*
 * Problem 			: Rotate Array
 * Topics           : Array, Math, Two Pointers
 * Level            : Medium
 * URL              : https://leetcode.com/problems/rotate-array
 * Description      : Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
 * 					  The rotation should be performed in-place with O(1) extra space.
 * Constraints		:
 * 					  - 1 <= nums.length <= 10^5
 * 					  - -2^31 <= nums[i] <= 2^31 - 1
 * 					  - 0 <= k <= 10^5
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,2,3,4,5,6,7], k = 3
 * 					  Output: [5,6,7,1,2,3,4]
 * 					  Explanation:
 * 					  rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * 					  rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * 					  rotate 3 steps to the right: [5,6,7,1,2,3,4]
 *
 * 					  Example 2:
 * 					  Input: nums = [-1,-100,3,99], k = 2
 * 					  Output: [3,99,-1,-100]
 * 					  Explanation:
 * 					  rotate 1 steps to the right: [99,-1,-100,3]
 * 					  rotate 2 steps to the right: [3,99,-1,-100]
 */

func rotateArray(nums []int, k int) {
	n := len(nums)

	if k%n == 0 {
		return
	}

	k = k % n

	// reverse all numbers
	reverseArray(nums, 0, n-1)
	// reverse first k numbers
	reverseArray(nums, 0, k-1)
	// reverse k -> n-1 numbers
	reverseArray(nums, k, n-1)
}

func reverseArray(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func RunTestRotateArray() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect []int
	}{
		"case-1": {
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:      4,
			expect: []int{5, 6, 7, 8, 1, 2, 3, 4},
		},
		"case-2": {
			nums:   []int{1000, 2, 4, -3},
			k:      2,
			expect: []int{4, -3, 1000, 2},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		rotateArray(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualSlices(testCase.nums, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, testCase.nums)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
