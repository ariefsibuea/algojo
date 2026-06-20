package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("MaximumSumOfDistinctSubarraysWithLengthK", RunTestMaximumSumOfDistinctSubarraysWithLengthK)
}

/*
 * Problem	: Maximum Sum of Distinct Subarrays With Length K
 * Topics	: Array, Hash Table, Sliding Window
 * Level	: Medium
 * URL		: https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
 *
 * Description:
 * 		Given an integer array nums and an integer k, find the maximum sum among all subarrays of length k where
 * 		all elements are distinct. A subarray is a contiguous non-empty sequence of elements within the array.
 * 		Return the maximum subarray sum that satisfies both conditions (length k and all distinct elements).
 * 		If no subarray of length k has all distinct elements, return 0.
 *
 * Constraints:
 * 		- 1 <= nums[i] <= 10^5
 * 		- 1 <= k <= nums.length <= 10^5
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [1,5,4,2,9,9,9], k = 3
 * 		Output: 15
 * 		Explanation: Subarrays of length 3 are: [1,5,4] (sum=10), [5,4,2] (sum=11), [4,2,9] (sum=15), [2,9,9] (invalid),
 * 		[9,9,9] (invalid). The maximum sum among valid subarrays is 15.
 *
 * 		Example 2:
 * 		Input: nums = [4,4,4], k = 3
 * 		Output: 0
 * 		Explanation: The only subarray of length 3 is [4,4,4] which has duplicate elements, so return 0.
 */

func maximumSubarraySum(nums []int, k int) int64 {
	sum, maxSum := int64(0), int64(0)
	seen := map[int]int{}

	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		seen[nums[i]]++
	}

	if len(seen) == k {
		maxSum = sum
	}

	for i := k; i < len(nums); i++ {
		sum = sum + int64(nums[i]) - int64(nums[i-k])
		seen[nums[i]]++
		seen[nums[i-k]]--
		if seen[nums[i-k]] == 0 {
			delete(seen, nums[i-k])
		}
		if len(seen) == k && maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum
}

func RunTestMaximumSumOfDistinctSubarraysWithLengthK() {
	runner.InitMetrics("MaximumSumOfDistinctSubarraysWithLengthK")

	testCases := map[string]struct {
		nums   []int
		k      int
		expect int64
	}{
		"example-1-basic": {
			nums:   []int{1, 5, 4, 2, 9, 9, 9},
			k:      3,
			expect: 15,
		},
		"example-2-all-duplicates": {
			nums:   []int{4, 4, 4},
			k:      3,
			expect: 0,
		},
		"all-distinct-max-window": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      3,
			expect: 12,
		},
		"mixed-duplicates-valid-window": {
			nums:   []int{1, 2, 1, 2, 3},
			k:      2,
			expect: 5,
		},
		"single-element": {
			nums:   []int{5},
			k:      1,
			expect: 5,
		},
		"k-equals-length-all-distinct": {
			nums:   []int{10, 20, 30},
			k:      3,
			expect: 60,
		},
		"k-equals-length-has-duplicates": {
			nums:   []int{10, 10, 10},
			k:      3,
			expect: 0,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"nums": tc.nums, "k": tc.k})

		result := runner.ExecCountMetrics(maximumSubarraySum, tc.nums, tc.k).(int64)
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
