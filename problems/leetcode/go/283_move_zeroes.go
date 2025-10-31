package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Move Zeroes
 * Topic            : Array, Two Pointers
 * Level            : Easy
 * URL              : https://leetcode.com/problems/move-zeroes
 * Description      : Given an integer array nums, move all 0's to the end of it while maintaining the relative order
 * 					of the non-zero elements. Note that you must do this in-place without making a copy of the array.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [0,1,0,3,12]
 * 					Output: [1,3,12,0,0]
 *
 * 					Example 2:
 * 					Input: nums = [0]
 * 					Output: [0]
 */

func moveZeroes(nums []int) {
	indexNonZero := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[indexNonZero], nums[i] = nums[i], nums[indexNonZero]
			indexNonZero += 1
		}
	}
}

func RunTestMoveZeroes() {
	testCases := map[string]struct {
		nums   []int
		expect []int
	}{
		"case-1": {
			nums:   []int{0, 1, 0, 3, 12},
			expect: []int{1, 3, 12, 0, 0},
		},
		"case-2": {
			nums:   []int{0},
			expect: []int{0},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		moveZeroes(testCase.nums)
		if !cmp.EqualSlices(testCase.nums, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, testCase.nums)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
