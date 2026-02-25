package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("TwoSumLessThanK", RunTestTwoSumLessThanK)
}

/**
 * Problem 			: Two Sum Less Than K
 * Topics           : Sort, Search
 * Level            : Easy
 * URL              : https://leetcode.com/problems/two-sum-less-than-k
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [2,1,3,3,5], k = 4
 * 					Output: 3
 * 					Explanation: The maximum sum less than k is 3, i.e., 2+1 = 3.
 *
 * 					Example 2:
 * 					Input: nums = [8,4,9,2,10,1], k = 9
 * 					Output: 6
 * 					Explanation: The maximum sum less than k is 6, i.e., 4+2 = 6.
 *
 * 					Example 3:
 * 					Input: nums = [3,4,5,8,6,2], k = 5
 * 					Output: -1
 *
 * 					Example 4:
 * 					Input: nums = [4,4,4,4,4,4,4,4], k = 12
 * 					Output: 8
 */

// twoSumLessThanK uses two pointers to solve the problem.
func twoSumLessThanK(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if nums[0] >= k {
		return -1
	}

	result := -1
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]+nums[right] >= k {
			right -= 1
		} else {
			result = max(result, nums[left]+nums[right])
			left += 1
		}
	}

	return result
}

// twoSumLessThanKSolutionI uses binary search to solve the problem.
func twoSumLessThanKSolutionI(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if nums[0] >= k {
		return -1
	}

	result := -1
	numsLen := len(nums)

	for i, num := range nums {
		j, t := -1, k-num
		left, right := i+1, numsLen-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] < t {
				j = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if j > i {
			result = max(result, nums[i]+nums[j])
		}
	}

	return result
}

func RunTestTwoSumLessThanK() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"case-1": {
			nums:   []int{2, 1, 3, 3, 5},
			k:      4,
			expect: 3,
		},
		"case-2": {
			nums:   []int{8, 4, 9, 2, 10, 1},
			k:      9,
			expect: 6,
		},
		"case-3": {
			nums:   []int{3, 4, 5, 8, 6, 2},
			k:      5,
			expect: -1,
		},
		"case-4": {
			nums:   []int{4, 4, 4, 4, 4, 4, 4, 4},
			k:      12,
			expect: 8,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := twoSumLessThanK(testCase.nums, testCase.k)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
