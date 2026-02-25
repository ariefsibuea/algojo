package main

import (
	cmp_ "cmp"
	"fmt"
	"os"
	"slices"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("FrequencyOfTheMostFrequentElement", RunTestFrequencyOfTheMostFrequentElement)
}

/*
 * Problem          : Frequency of the Most Frequent Element
 * Topics           : Array, Binary Search, Greedy, Sliding Window, Sorting, Prefix Sum
 * Level            : Medium
 * URL              : https://leetcode.com/problems/frequency-of-the-most-frequent-element
 * Description      : The frequency of an element is the number of times it occurs in an array. You are given an
 *                    integer array nums and an integer k. In one operation, you can choose an index of nums and
 *                    increment the element at that index by 1. Return the maximum possible frequency of an element
 *                    after performing at most k operations.
 * Constraints      :
 *                    - 1 <= nums.length <= 10^5
 *                    - 1 <= nums[i] <= 10^5
 *                    - 1 <= k <= 10^5
 * Examples         :
 *                    Example 1:
 *                    Input: nums = [1,2,4], k = 5
 *                    Output: 3
 *                    Explanation: Increment the first element three times and the second element two times to make
 *                    nums = [4,4,4]. 4 has a frequency of 3.
 *
 *                    Example 2:
 *                    Input: nums = [1,4,8,13], k = 5
 *                    Output: 2
 *                    Explanation: There are multiple optimal solutions:
 *                    - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
 *                    - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
 *                    - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.
 *
 *                    Example 3:
 *                    Input: nums = [3,9,6], k = 2
 *                    Output: 1
 */

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	sortFunc := func(i, j int) int {
		return cmp_.Compare(i, j)
	}
	slices.SortFunc(nums, sortFunc)

	maxFreq := 0
	sum := 0
	start := 0

	for end := 0; end < n; end++ {
		sum += nums[end]

		if ((end-start+1)*nums[end])-sum > k && start < end {
			sum = sum - nums[start]
			start++
		}

		maxFreq = max(maxFreq, end-start+1)
	}

	return maxFreq
}

func RunTestFrequencyOfTheMostFrequentElement() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"example-1": {
			nums:   []int{1, 2, 4},
			k:      5,
			expect: 3,
		},
		"example-2": {
			nums:   []int{1, 4, 8, 13},
			k:      5,
			expect: 2,
		},
		"example-3": {
			nums:   []int{3, 9, 6},
			k:      2,
			expect: 1,
		},
		"single-element": {
			nums:   []int{5},
			k:      10,
			expect: 1,
		},
		"all-elements-same": {
			nums:   []int{5, 5, 5, 5},
			k:      3,
			expect: 4,
		},
		"k-zero": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      0,
			expect: 1,
		},
		"large-k": {
			nums:   []int{1, 2, 3},
			k:      100000,
			expect: 3,
		},
		"two-elements-k-enough": {
			nums:   []int{1, 5},
			k:      4,
			expect: 2,
		},
		"two-elements-k-not-enough": {
			nums:   []int{1, 10},
			k:      5,
			expect: 1,
		},
		"already-sorted": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      3,
			expect: 3,
		},
		"reverse-sorted": {
			nums:   []int{5, 4, 3, 2, 1},
			k:      3,
			expect: 3,
		},
		"multiple-same-elements": {
			nums:   []int{1, 1, 2, 2, 3, 3},
			k:      2,
			expect: 4,
		},
		"consecutive-numbers": {
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      6,
			expect: 4,
		},
		"large-numbers": {
			nums:   []int{100000, 99999, 99998},
			k:      2,
			expect: 2,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := maxFrequency(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
