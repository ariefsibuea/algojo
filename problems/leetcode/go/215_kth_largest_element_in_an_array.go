package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	// register function here
}

/*
 * Problem	: 215. Kth Largest Element in an Array
 * Topics	: Array, Divide and Conquer, Sorting, Heap (Priority Queue), Quickselect
 * Level	: Medium
 * URL		: https://leetcode.com/problems/kth-largest-element-in-an-array/
 *
 * Description:
 *		Given an integer array 'nums' and an integer 'k', return the kth largest element in the array. Note that it is
 *		the kth largest element in the sorted order, not the kth distinct element. The task can be solved using various
 *		approaches including sorting, using a min-heap of size k, or applying the Quickselect algorithm for optimal
 *		average time complexity.
 *
 * Constraints:
 *		- 1 <= k <= nums.length <= 10^5
 *		- -10^4 <= nums[i] <= 10^4
 *
 * Examples:
 *		Example 1:
 *		Input: nums = [3,2,1,5,6,4], k = 2
 *		Output: 5
 *		Explanation: The sorted array in descending order is [6,5,4,3,2,1]. The 2nd largest element is 5.
 *
 *		Example 2:
 *		Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 *		Output: 4
 *		Explanation: The sorted array in descending order is [6,5,5,4,3,3,2,2,1]. The 4th largest element is 4.
 */

func findKthLargest(nums []int, k int) int {
	return 0
}

func RunTestKthLargestElementInAnArray() {
	runner.InitMetrics("KthLargestElementInAnArray")

	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"example-1-basic-case": {
			nums:   []int{3, 2, 1, 5, 6, 4},
			k:      2,
			expect: 5,
		},
		"example-2-duplicates": {
			nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			expect: 4,
		},
		"single-element": {
			nums:   []int{1},
			k:      1,
			expect: 1,
		},
		"all-same-elements": {
			nums:   []int{5, 5, 5, 5},
			k:      2,
			expect: 5,
		},
		"negative-numbers": {
			nums:   []int{-1, -2, -3, -4},
			k:      2,
			expect: -2,
		},
		"k-equals-1-largest": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      1,
			expect: 5,
		},
		"k-equals-n-smallest": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      5,
			expect: 1,
		},
		"unsorted-large-array": {
			nums:   []int{10, 3, 8, 1, 7, 2, 9, 4, 6, 5},
			k:      3,
			expect: 8,
		},
		"mixed-positive-negative": {
			nums:   []int{-5, 3, 0, -2, 4, -1},
			k:      3,
			expect: 1,
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"nums": testCase.nums, "k": testCase.k})

		// Create a copy of nums since the solution might modify it
		numsCopy := make([]int, len(testCase.nums))
		copy(numsCopy, testCase.nums)

		result := runner.ExecCountMetrics(findKthLargest, numsCopy, testCase.k).(int)
		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
