package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ValidTriangleNumber", RunTestValidTriangleNumber)
}

/*
 * Problem	: Valid Triangle Number
 * Topics	: Array, Two Pointers, Binary Search, Greedy, Sorting
 * Level	: Medium
 * URL		: https://leetcode.com/problems/valid-triangle-number/
 *
 * Description:
 * 		Given an integer array nums, return the number of triplets chosen from the array that can make
 * 		triangles if we take them as side lengths of a triangle.
 *
 * Constraints:
 * 		- 1 <= nums.length <= 1000
 * 		- 0 <= nums[i] <= 1000
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [2,2,3,4]
 * 		Output: 3
 * 		Explanation: Valid combinations are:
 * 		2,3,4 (using the first 2)
 * 		2,3,4 (using the second 2)
 * 		2,2,3
 *
 * 		Example 2:
 * 		Input: nums = [4,2,3,4]
 * 		Output: 4
 */

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	count := 0

	for i := len - 1; i > 1; i-- {
		n := nums[i]
		left := 0
		right := i - 1

		for left < right {
			// NOTE: A triangle is valid if a + b > c, b + c > a, a + c > b.
			if nums[left]+nums[right] > n {
				count += right - left
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return count
}

func RunTestValidTriangleNumber() {
	runner.InitMetrics("ValidTriangleNumber")

	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"basic": {
			nums:   []int{2, 2, 3, 4},
			expect: 3,
		},
		"duplicate-larger": {
			nums:   []int{4, 2, 3, 4},
			expect: 4,
		},
		"no-valid-triangle": {
			nums:   []int{1, 2, 3},
			expect: 0,
		},
		"all-equal-sides": {
			nums:   []int{3, 3, 3},
			expect: 1,
		},
		"large-array": {
			nums:   []int{2, 2, 3, 4, 5},
			expect: 6,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums})

		result := runner.ExecCountMetrics(triangleNumber, tc.nums).(int)
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
