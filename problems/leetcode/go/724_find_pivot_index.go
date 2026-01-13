package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Find Pivot Index
 * Topics           : Array, Prefix Sum
 * Level            : Easy
 * URL              : https://leetcode.com/problems/find-pivot-index
 * Description      : Given an array of integers `nums`, the task is to calculate the pivot index. A pivot index is
 *                    defined as the index where the sum of all numbers strictly to its left equals the sum of all
 *                    numbers strictly to its right. If the index is at the left edge, the left sum is 0, and
 *                    similarly for the right edge, the right sum is 0. The function should return the leftmost
 *                    pivot index, or -1 if no such index exists.
 * Examples         :
 *                    Example 1:
 *                    Input: nums = [1,7,3,6,5,6]
 *                    Output: 3
 *                    Explanation: The pivot index is 3. Left sum (nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11)
 *                    equals Right sum (nums[4] + nums[5] = 5 + 6 = 11).
 *
 *                    Example 2:
 *                    Input: nums = [1,2,3]
 *                    Output: -1
 *                    Explanation: No index satisfies the conditions.
 *
 *                    Example 3:
 *                    Input: nums = [2,1,-1]
 *                    Output: 0
 *                    Explanation: The pivot index is 0. Left sum is 0 (no elements to the left). Right sum
 *                    (nums[1] + nums[2] = 1 + -1 = 0) is also 0.
 */

func pivotIndex(nums []int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = nums[i] + prefixSum[i]
	}

	result := -1

	totalSum := prefixSum[n]
	for i := n; i >= 1; i-- {
		rightVal := totalSum - prefixSum[i]
		leftVal := prefixSum[i-1]
		if rightVal == leftVal {
			result = i - 1
		}
	}

	return result
}

func RunTestFindPivotIndex() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{1, 7, 3, 6, 5, 6},
			expect: 3,
		},
		"case-2": {
			nums:   []int{1, 2, 3},
			expect: -1,
		},
		"case-3": {
			nums:   []int{2, 1, -1},
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := pivotIndex(testCase.nums)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
