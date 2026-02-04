package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Subarrays with K Different Integers
 * Topics           : Array, Hash Table, Sliding Window
 * Level            : Hard
 * URL              : https://leetcode.com/problems/subarrays-with-k-different-integers
 * Description      : Given an integer array nums and an integer k, return the total number of subarrays that contain
 * 					  exactly k distinct integers. A subarray is a contiguous part of an array where the elements are
 * 					  consecutive.
 * Constraints      :
 * 					  - 1 <= nums.length <= 10^5
 * 					  - 1 <= k <= nums.length
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,2,1,2,3], k = 2
 * 					  Output: 7
 * 					  Explanation: The subarrays are [1,2], [2,1], [1,2,1], [2,1,2], [1,2,3], [2,1,2,3], [2,1,2,3]
 * 					  where each contains exactly 2 distinct integers.
 *
 * 					  Example 2:
 * 					  Input: nums = [1,2,1,2,3], k = 3
 * 					  Output: 3
 * 					  Explanation: The subarrays with exactly 3 distinct integers are [1,2,1,2], [2,1,2,3], [1,2,1,2,3]
 */

func subarraysWithKDistinct(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
	existed := make(map[int]int)
	count := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		existed[nums[end]]++

		for len(existed) > k {
			existed[nums[start]]--
			if existed[nums[start]] <= 0 {
				delete(existed, nums[start])
			}
			start++
		}

		count = count + (end - start + 1)
	}

	return count
}

func RunTestSubarraysWithKDifferentIntegers() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"example-1": {
			nums:   []int{1, 2, 1, 2, 3},
			k:      2,
			expect: 7,
		},
		"example-2": {
			nums:   []int{1, 2, 1, 2, 3},
			k:      3,
			expect: 3,
		},
		"single-element-k-1": {
			nums:   []int{1},
			k:      1,
			expect: 1,
		},
		"all-same-elements": {
			nums:   []int{1, 1, 1, 1},
			k:      1,
			expect: 10,
		},
		"all-different-elements": {
			nums:   []int{1, 2, 3},
			k:      3,
			expect: 1,
		},
		"k-equals-array-length": {
			nums:   []int{1, 2, 3},
			k:      3,
			expect: 1,
		},
		"with-negative-numbers": {
			nums:   []int{-1, -2, -1, -3},
			k:      2,
			expect: 4,
		},
		"k-1-minimum": {
			nums:   []int{1, 2, 3, 4},
			k:      1,
			expect: 4,
		},
		"duplicate-with-different-order": {
			nums:   []int{1, 1, 2, 2, 3},
			k:      2,
			expect: 6,
		},
		"two-elements": {
			nums:   []int{1, 2},
			k:      2,
			expect: 1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := subarraysWithKDistinct(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
