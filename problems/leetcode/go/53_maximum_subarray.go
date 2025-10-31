package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Maximum Subarray
 * Topics           : Array, Divide and Conquer, Dynamic Programming
 * Level            : Medium
 * URL              : https://leetcode.com/problems/maximum-subarray
 * Description      : Given an integer array nums, find the subarray with the largest sum, and return its sum.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 					Output: 6
 * 					Explanation: The subarray [4,-1,2,1] has the largest sum 6.
 *
 * 					Example 2:
 * 					Input: nums = [1]
 * 					Output: 1
 * 					Explanation: The subarray [1] has the largest sum 1.
 *
 * 					Example 3:
 * 					Input: nums = [5,4,-1,7,8]
 * 					Output: 23
 * 					Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
 */

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sumResult := 0
	maxSum := math.MinInt

	for _, num := range nums {
		if sumResult < 0 {
			sumResult = 0
		}
		sumResult += num
		maxSum = max(maxSum, sumResult)
	}

	return maxSum
}

func RunTestMaxSubArray() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
		"case-2": {
			nums:   []int{1},
			expect: 1,
		},
		"case-3": {
			nums:   []int{5, 4, -1, 7, 8},
			expect: 23,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := maxSubArray(testCase.nums)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
