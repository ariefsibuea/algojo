package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Two Sum
 * Topics           : Array, Hash Table
 * Level            : Easy
 * URL              : https://leetcode.com/problems/two-sum
 * Description      : Given an array of integers nums and an integer target, return indices of the two numbers such
 * 					that they add up to target. You may assume that each input would have exactly one solution, and
 * 					you may not use the same element twice. You can return the answer in any order.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [2,7,11,15], target = 9
 * 					Output: [0,1]
 * 					Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 *
 * 					Example 2:
 * 					Input: nums = [3,2,4], target = 6
 * 					Output: [1,2]
 *
 * 					Example 3:
 * 					Input: nums = [3,3], target = 6
 * 					Output: [0,1]
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

	var result []int

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result = twoSum(testCase.nums, testCase.target)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
