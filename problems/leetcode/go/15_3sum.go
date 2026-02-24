package main

import (
	"fmt"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem	: 3Sum
 * Topics	: Array, Two Pointers, Sorting
 * Level	: Medium
 * URL		: https://leetcode.com/problems/3sum/
 *
 * Description:
 * 		You are given an integer array nums. Your task is to find all unique triplets in the array where
 * 		three numbers add up to zero. Specifically, you need to return all triplets [nums[i], nums[j],
 * 		nums[k]] that satisfy these conditions:
 * 		- The three indices must be different: i != j, i != k, and j != k
 * 		- The sum equals zero: nums[i] + nums[j] + nums[k] == 0
 * 		- The solution set must not contain duplicate triplets (even if the same values appear
 * 		multiple times in the array)
 *
 * Constraints:
 * 		- 3 <= nums.length <= 3000
 * 		- -10^5 <= nums[i] <= 10^5
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [-1,0,1,2,-1,-4]
 * 		Output: [[-1,-1,2],[-1,0,1]]
 * 		Explanation:
 * 		nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * 		nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * 		nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * 		The distinct triplets are [-1,0,1] and [-1,-1,2].
 * 		Notice that the order of the output and the order of the triplets does not matter.
 *
 * 		Example 2:
 * 		Input: nums = [0,1,1]
 * 		Output: []
 * 		Explanation: The only possible triplet does not sum up to 0.
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	len := len(nums)
	result := make([][]int, 0)

	for i := 0; i < len-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len - 1

		for j < k {
			switch sum := nums[i] + nums[j] + nums[k]; {
			case sum > 0:
				k -= 1
			case sum < 0:
				j += 1
			default:
				result = append(result, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j += 1
				}
				j += 1
			}
		}
	}

	return result
}

func RunTestThreeSum() {
	runner.InitMetrics("ThreeSum")

	testCases := map[string]struct {
		nums   []int
		expect [][]int
	}{
		"example-1-basic": {
			nums: []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		"example-2-no-triplet": {
			nums:   []int{0, 1, 1},
			expect: [][]int{},
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums})

		result := runner.ExecCountMetrics(threeSum, tc.nums).([][]int)
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
