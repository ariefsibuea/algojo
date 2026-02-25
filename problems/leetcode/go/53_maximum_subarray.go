package main

import (
	"fmt"
	"math"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("MaximumSubarray", RunTestMaximumSubarray)
}

/*
 * Problem	: Maximum Subarray
 * Topics	: Array, Divide and Conquer, Dynamic Programming
 * Level	: Medium
 * URL		: https://leetcode.com/problems/maximum-subarray/
 *
 * Description:
 * 		Given an integer array nums, find the subarray with the largest sum, and return its sum.
 *
 * Constraints:
 * 		- 1 <= nums.length <= 10^5
 * 		- -10^4 <= nums[i] <= 10^4
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 		Output: 6
 * 		Explanation: The subarray [4,-1,2,1] has the largest sum 6.
 *
 * 		Example 2:
 * 		Input: nums = [1]
 * 		Output: 1
 * 		Explanation: The subarray [1] has the largest sum 1.
 *
 * 		Example 3:
 * 		Input: nums = [5,4,-1,7,8]
 * 		Output: 23
 * 		Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
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

func RunTestMaximumSubarray() {
	runner.InitMetrics("MaximumSubarray")

	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"example-1-basic": {
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
		"example-2-single-element": {
			nums:   []int{1},
			expect: 1,
		},
		"example-3-all-positive": {
			nums:   []int{5, 4, -1, 7, 8},
			expect: 23,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums})

		result := runner.ExecCountMetrics(maxSubArray, tc.nums).(int)
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
