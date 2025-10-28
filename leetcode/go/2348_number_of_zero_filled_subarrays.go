package main

import (
	"fmt"
	"os"
)

/**
 * Problem 			: Number of Zero Filled Subarrays
 * Topics           : Array, Math
 * Level            : Medium
 * URL              : https://leetcode.com/problems/number-of-zero-filled-subarrays
 * Description      : Given an integer array nums, return the number of subarrays filled with 0.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [1,3,0,0,2,0,0,4]
 * 					Output: 6
 * 					Explanation:
 * 					There are 4 occurrences of [0] as a subarray.
 * 					There are 2 occurrences of [0,0] as a subarray.
 * 					There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.
 *
 * 					Example 2:
 * 					Input: nums = [0,0,0,2,0,0]
 * 					Output: 9
 * 					Explanation:
 * 					There are 5 occurrences of [0] as a subarray.
 * 					There are 3 occurrences of [0,0] as a subarray.
 * 					There is 1 occurrence of [0,0,0] as a subarray.
 * 					There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.
 *
 * 					Example 3:
 * 					Input: nums = [2,10,2019]
 * 					Output: 0
 * 					Explanation: There is no subarray filled with 0. Therefore, we return 0.
 */

func zeroFilledSubarray(nums []int) int64 {
	consecutiveZero := 0
	totalSubarray := int64(0)

	for _, num := range nums {
		if num == 0 {
			consecutiveZero += 1
			totalSubarray += int64(consecutiveZero)
		} else {
			consecutiveZero = 0
		}
	}

	return totalSubarray
}

func RunTestZeroFilledSubarray() {
	testCases := map[string]struct {
		nums   []int
		expect int64
	}{
		"case-1": {
			nums:   []int{1, 3, 0, 0, 2, 0, 0, 4},
			expect: 6,
		},
		"case-2": {
			nums:   []int{0, 0, 0, 2, 0, 0},
			expect: 9,
		},
		"case-3": {
			nums:   []int{2, 10, 2019},
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := zeroFilledSubarray(testCase.nums)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
