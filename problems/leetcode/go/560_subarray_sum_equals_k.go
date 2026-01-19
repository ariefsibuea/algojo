package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Subarray Sum Equals K
 * Topics           : Array, Hash Table, Prefix Sum
 * Level            : Medium
 * URL              : https://leetcode.com/problems/subarray-sum-equals-k
 * Description      : Given an array of integers nums and an integer k, return the total number of subarrays whose sum
 * 					  equals to k. A subarray is a contiguous non-empty sequence of elements within an array.
 *
 * 					  Constraints:
 * 					  - 1 <= nums.length <= 2 * 10^4
 * 					  - -1000 <= nums[i] <= 1000
 * 					  - -10^7 <= k <= 10^7
 *
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,1,1], k = 2
 * 					  Output: 2
 *
 * 					  Example 2:
 * 					  Input: nums = [1,2,3], k = 3
 * 					  Output: 2
 */

func subarraySum(nums []int, k int) int {
	prefixSum := map[int]int{0: 1}

	result, sum := 0, 0

	for _, num := range nums {
		sum += num
		dif := sum - k
		result += prefixSum[dif]
		prefixSum[sum]++
	}

	return result
}

func RunTestSubarraySumEqualsK() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect int
	}{
		"case-1": {
			nums:   []int{1, 1, 1},
			k:      2,
			expect: 2,
		},
		"case-2": {
			nums:   []int{1, 2, 3},
			k:      3,
			expect: 2,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := subarraySum(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums-1": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
