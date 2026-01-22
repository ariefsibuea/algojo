package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Max Consecutive Ones III
 * Topics           : Array, Binary Search, Sliding Window, Prefix Sum
 * Level            : Medium
 * URL              : https://leetcode.com/problems/max-consecutive-ones-iii
 * Description      : Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the
 * 					  array if you can flip at most k 0's.
 * Constraints		:
 * 					  - 1 <= nums.length <= 10^5
 * 					  - nums[i] is either 0 or 1
 * 					  - 0 <= k <= nums.length
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
 * 					  Output: 6
 * 					  Explanation: The longest subarray with at most 2 zeros is [0,0,1,1,1,1] (from index 4 to 9),
 * 					  which has length 6. Flipping the two 0s results in all 1s.
 *
 * 					  Example 2:
 * 					  Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
 * 					  Output: 10
 * 					  Explanation: The longest subarray with at most 3 zeros is [0,0,1,1,0,0,1,1,1,0,1,1]
 * 					  (from index 2 to 11), which has length 10.
 */

func longestOnes(nums []int, k int) int {
	flip := 0
	maxLength := 0

	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			flip++
		}

		for flip > k && start <= end {
			if nums[start] == 0 {
				flip--
			}
			start++
		}

		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func RunTestMaxConsecutiveOnesIII() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"example-1": {
			nums:   []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:      2,
			expect: 6,
		},
		"example-2": {
			nums:   []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:      3,
			expect: 10,
		},
		"all-ones": {
			nums:   []int{1, 1, 1, 1, 1},
			k:      2,
			expect: 5,
		},
		"all-zeros-with-flips": {
			nums:   []int{0, 0, 0, 0, 0},
			k:      3,
			expect: 3,
		},
		"all-zeros-not-enough-flips": {
			nums:   []int{0, 0, 0, 0, 0},
			k:      2,
			expect: 2,
		},
		"k-is-zero": {
			nums:   []int{1, 1, 0, 1, 1, 0, 1},
			k:      0,
			expect: 2,
		},
		"k-is-equal-to-length": {
			nums:   []int{0, 0, 0, 1, 1, 0},
			k:      6,
			expect: 6,
		},
		"k-larger-than-zeros": {
			nums:   []int{1, 0, 1, 0, 1},
			k:      3,
			expect: 5,
		},
		"zeros-at-beginning": {
			nums:   []int{0, 0, 1, 1, 1},
			k:      2,
			expect: 5,
		},
		"zeros-at-end": {
			nums:   []int{1, 1, 1, 0, 0},
			k:      2,
			expect: 5,
		},
		"single-one": {
			nums:   []int{1},
			k:      1,
			expect: 1,
		},
		"single-zero-can-flip": {
			nums:   []int{0},
			k:      1,
			expect: 1,
		},
		"single-zero-cannot-flip": {
			nums:   []int{0},
			k:      0,
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := longestOnes(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
