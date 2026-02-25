package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("CountSubarraysWithScoreLessThanK", RunTestCountSubarraysWithScoreLessThanK)
}

/*
 * Problem 			: Count Subarrays With Score Less Than K
 * Topics           : Array, Sliding Window
 * Level            : Hard
 * URL              : https://leetcode.com/problems/count-subarrays-with-score-less-than-k
 * Description      : The score of an array is defined as the product of its sum and its length. Given a positive
 *                    integer array nums and an integer k, return the number of non-empty subarrays of nums whose
 *                    score is strictly less than k. A subarray is a contiguous sequence of elements within an array.
 * Constraints      :
 *                    - 1 <= nums.length <= 10^5
 *                    - 1 <= nums[i] <= 10^5
 *                    - 1 <= k <= 10^10
 * Examples         :
 *                    Example 1:
 *                    Input: nums = [2,1,4,3,5], k = 10
 *                    Output: 6
 *                    Explanation: The 6 subarrays having scores less than 10 are:
 *                    - [2] with score 2 * 1 = 2
 *                    - [1] with score 1 * 1 = 1
 *                    - [4] with score 4 * 1 = 4
 *                    - [3] with score 3 * 1 = 3
 *                    - [5] with score 5 * 1 = 5
 *                    - [2,1] with score (2 + 1) * 2 = 6
 *
 *                    Example 2:
 *                    Input: nums = [1,1,1], k = 5
 *                    Output: 5
 *                    Explanation: Every subarray except [1,1,1] has a score less than 5. The subarray [1,1,1] has a
 *                    score of (1 + 1 + 1) * 3 = 9 which is greater than 5.
 */

func countSubarrays(nums []int, k int64) int64 {
	var count int64 = 0
	var sum, score int64 = 0, 0
	var start = 0

	for end := 0; end < len(nums); end++ {
		sum += int64(nums[end])
		score = sum * int64(end-start+1)

		for score >= k {
			sum -= int64(nums[start])
			start++
			score = sum * int64(end-start+1)
		}

		count += int64(end - start + 1)
	}

	return count
}

func RunTestCountSubarraysWithScoreLessThanK() {
	testCases := map[string]struct {
		nums   []int
		k      int64
		expect int64
	}{
		"example-1": {
			nums:   []int{2, 1, 4, 3, 5},
			k:      10,
			expect: 6,
		},
		"example-2": {
			nums:   []int{1, 1, 1},
			k:      5,
			expect: 5,
		},
		"single-element-valid": {
			nums:   []int{1},
			k:      2,
			expect: 1,
		},
		"single-element-invalid": {
			nums:   []int{5},
			k:      5,
			expect: 0,
		},
		"all-elements-too-large": {
			nums:   []int{100000, 100000},
			k:      100000,
			expect: 0,
		},
		"large-k-all-subarrays-valid": {
			nums:   []int{1, 2, 3},
			k:      100,
			expect: 6,
		},
		"empty-subarrays-only": {
			nums:   []int{10, 10, 10},
			k:      10,
			expect: 0,
		},
		"two-elements-one-valid": {
			nums:   []int{1, 10},
			k:      10,
			expect: 1,
		},
		"all-ones-small-k": {
			nums:   []int{1, 1, 1, 1},
			k:      3,
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := countSubarrays(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
