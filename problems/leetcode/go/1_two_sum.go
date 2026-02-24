package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: Two Sum
 * Topics	: Array, Hash Table
 * Level	: Easy
 * URL		: https://leetcode.com/problems/two-sum/
 *
 * Description:
 * 		Given an array of integers nums and an integer target, return indices of the two numbers such
 * 		that they add up to target. You may assume that each input would have exactly one solution, and
 * 		you may not use the same element twice. You can return the answer in any order.
 *
 * Constraints:
 * 		- 2 <= nums.length <= 10^4
 * 		- -10^9 <= nums[i] <= 10^9
 * 		- -10^9 <= target <= 10^9
 * 		- Only one valid answer exists.
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [2,7,11,15], target = 9
 * 		Output: [0,1]
 * 		Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 *
 * 		Example 2:
 * 		Input: nums = [3,2,4], target = 6
 * 		Output: [1,2]
 *
 * 		Example 3:
 * 		Input: nums = [3,3], target = 6
 * 		Output: [0,1]
 */

func twoSum(nums []int, target int) []int {
	numExists := make(map[int]int)

	for i, n := range nums {
		remaining := target - n
		if idx, ok := numExists[remaining]; ok {
			return []int{idx, i}
		}
		numExists[n] = i
	}

	return []int{}
}

func RunTestTwoSum() {
	runner.InitMetrics("TwoSum")

	testCases := map[string]struct {
		nums   []int
		target int
		expect []int
	}{
		"case-1": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		"case-2": {
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		"case-3": {
			nums:   []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums, "target": tc.target})

		result := runner.ExecCountMetrics(twoSum, tc.nums, tc.target).([]int)
		if !cmp.EqualSlices(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
