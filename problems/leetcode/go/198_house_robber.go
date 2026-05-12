package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("HouseRobber", RunTestHouseRobber)
}

/*
 * Problem	: House Robber
 * Topics	: Array, Dynamic Programming
 * Level	: Medium
 * URL		: https://leetcode.com/problems/house-robber/
 *
 * Description:
 * 		You are a professional robber planning to rob houses along a street. Each house holds a certain amount of
 * 		money, but adjacent houses are connected by a security system that automatically alerts the police if two
 * 		neighbouring houses are broken into on the same night. Given an integer array nums where each element
 * 		represents the cash stashed in the i-th house, determine the maximum amount of money you can steal tonight
 * 		without triggering the alarm.
 *
 * Constraints:
 * 		- 1 <= nums.length <= 100
 * 		- 0 <= nums[i] <= 400
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [1,2,3,1]
 * 		Output: 4
 * 		Explanation: Rob house 1 (money = 1) and house 3 (money = 3) for a total of 1 + 3 = 4. House 2 is skipped
 * 		because it is adjacent to house 1, and house 4 is skipped because it is adjacent to house 3.
 *
 * 		Example 2:
 * 		Input: nums = [2,7,9,3,1]
 * 		Output: 12
 * 		Explanation: Rob house 1 (money = 2), house 3 (money = 9), and house 5 (money = 1) for a total of
 * 		2 + 9 + 1 = 12. Adjacent houses are never robbed on the same night.
 */

func robMemoization(nums []int) int {
	maxMoney := make([]int, len(nums))
	for i := range nums {
		maxMoney[i] = -1
	}

	var robHouse func(i int) int
	robHouse = func(i int) int {
		if i < 0 {
			return 0
		}
		if maxMoney[i] != -1 {
			return maxMoney[i]
		}

		maxMoney[i] = max(robHouse(i-1), robHouse(i-2)+nums[i])
		return maxMoney[i]
	}

	return robHouse(len(nums) - 1)
}

func robTabulation(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxMoney := make([]int, len(nums))
	maxMoney[0] = nums[0]
	maxMoney[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		maxMoney[i] = max(maxMoney[i-1], maxMoney[i-2]+nums[i])
	}

	return maxMoney[len(nums)-1]
}

func RunTestHouseRobber() {
	runner.InitMetrics("HouseRobber")

	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"example-1-basic": {
			nums:   []int{1, 2, 3, 1},
			expect: 4,
		},
		"example-2-alternating": {
			nums:   []int{2, 7, 9, 3, 1},
			expect: 12,
		},
		"single-house": {
			nums:   []int{5},
			expect: 5,
		},
		"two-houses-choose-max": {
			nums:   []int{3, 1},
			expect: 3,
		},
		"all-zeros": {
			nums:   []int{0, 0, 0, 0},
			expect: 0,
		},
		"large-values": {
			nums:   []int{400, 0, 400},
			expect: 800,
		},
		"increasing-values": {
			nums:   []int{1, 2, 3, 4, 5},
			expect: 9,
		},
	}

	var passedCount int

	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"nums": tc.nums})

		result := runner.ExecCountMetrics(robMemoization, tc.nums).(int)
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
