package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Minimum Size Subarray Sum
 * Topics           : Array, Binary Search, Sliding Window, Prefix Sum
 * Level            : Medium
 * URL              : https://leetcode.com/problems/minimum-size-subarray-sum
 * Description      : Given an array of positive integers `nums` and a positive integer `target`, return the minimal
 * 					  length of a subarray whose sum is greater than or equal to `target`. If there is no such
 * 					  subarray, return 0 instead.
 * Constraints      :
 *                    - 1 <= target <= 10^9
 *                    - 1 <= nums.length <= 10^5
 *                    - 1 <= nums[i] <= 10^4
 * Examples         :
 *                    Example 1:
 *                    Input: target = 7, nums = [2,3,1,2,4,3]
 *                    Output: 2
 *                    Explanation: The subarray [4,3] has the minimal length under the problem constraint.
 *
 *                    Example 2:
 *                    Input: target = 4, nums = [1,4,4]
 *                    Output: 1
 *
 *                    Example 3:
 *                    Input: target = 11, nums = [1,1,1,1,1,1,1,1]
 *                    Output: 0
 */

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := math.MaxInt
	sum := 0

	start := 0
	for end := 0; end < n; end++ {
		sum += nums[end]

		for sum >= target && start <= end {
			minLength = min(minLength, end-start+1)

			sum -= nums[start]
			start++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}

func RunTestMinimumSizeSubarraySum() {
	testCases := map[string]struct {
		target int
		nums   []int
		expect int
	}{
		"basic-case": {
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			expect: 2,
		},
		"no-solution": {
			target: 20,
			nums:   []int{1, 1, 1, 1},
			expect: 0,
		},
		"target-is-one-element": {
			target: 4,
			nums:   []int{1, 4, 4},
			expect: 1,
		},
		"all-elements-needed": {
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			expect: 5,
		},
		"one-element-array-match": {
			target: 5,
			nums:   []int{5},
			expect: 1,
		},
		"one-element-array-no-match": {
			target: 5,
			nums:   []int{4},
			expect: 0,
		},
		"empty-array": {
			target: 100,
			nums:   []int{},
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := minSubArrayLen(testCase.target, testCase.nums)
		format.PrintInput(map[string]interface{}{"target": testCase.target, "nums": testCase.nums})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
