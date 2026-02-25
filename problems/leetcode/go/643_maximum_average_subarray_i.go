package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("MaximumAverageSubarrayI", RunTestMaximumAverageSubarrayI)
}

/*
 * Problem 			: Maximum Average Subarray I
 * Topics           : Array, Sliding Window
 * Level            : Easy
 * URL              : https://leetcode.com/problems/maximum-average-subarray-i/
 * Description      : You are given an integer array nums consisting of n elements, and an integer k. Find a contiguous
 *                    subarray whose length is equal to k that has the maximum average value and return this value.
 *                    Any answer with a calculation error less than 10^-5 will be accepted.
 * Constraints		:
 * 					  - n == nums.length
 * 					  - 1 <= k <= n <= 10^5
 * 					  - -10^4 <= nums[i] <= 10^4
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,12,-5,-6,50,3], k = 4
 * 					  Output: 12.75000
 * 					  Explanation: The subarray [12, -5, -6, 50] has the maximum
 * 					  average of (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75.
 *
 * 					  Example 2:
 * 					  Input: nums = [5], k = 1
 * 					  Output: 5.00000
 * 					  Explanation: The subarray [5] has a sum of 5 and a length of 1, so the average is 5.0.
 */

func findMaxAverage(nums []int, k int) float64 {
	var sum, maxSum = 0, 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum = sum
	start := 0
	for end := k; end < len(nums); end++ {
		sum = sum + nums[end] - nums[start]
		maxSum = max(maxSum, sum)
		start++
	}

	return float64(maxSum) / float64(k)
}

func RunTestMaximumAverageSubarrayI() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect float64
	}{
		"leetcode-example-1": {
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			expect: 12.75,
		},
		"leetcode-example-2": {
			nums:   []int{5},
			k:      1,
			expect: 5.0,
		},
		"all-negative-numbers": {
			nums:   []int{-1, -2, -3, -4},
			k:      2,
			expect: -1.5,
		},
		"k-equals-length-of-nums": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      5,
			expect: 3.0,
		},
		"array-with-zeros": {
			nums:   []int{0, 4, 0, 3, 0},
			k:      2,
			expect: 2.0,
		},
		"maximum-average-at-the-beginning": {
			nums:   []int{10, 2, 3, 1, 5},
			k:      2,
			expect: 6.0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := findMaxAverage(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
