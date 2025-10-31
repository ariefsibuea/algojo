package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Longest Increasing Subsequence
 * Topics           : Array, Binary Search, Dynamic Programming
 * Level            : Medium
 * URL              : https://leetcode.com/problems/longest-increasing-subsequence
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [10,9,2,5,3,7,101,18]
 * 					Output: 4
 * 					Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 *
 * 					Example 2:
 * 					Input: nums = [0,1,0,3,2,3]
 * 					Output: 4
 *
 * 					Example 3:
 * 					Input: nums = [7,7,7,7,7,7,7]
 * 					Output: 1
 */

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}

	longest := 1

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
		longest = max(longest, lis[i])
	}

	return longest
}

func RunTestLongestIncreasingSubsequence() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		},
		"case-2": {
			nums:   []int{0, 1, 0, 3, 2, 3},
			expect: 4,
		},
		"case-3": {
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			expect: 1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := lengthOfLIS(testCase.nums)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
