package main

import (
	"fmt"
	"math"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem			: Median of Two Sorted Arrays
 * Topics			: Array, Binary Search, Divide and Conquer
 * Level			: Hard
 * URL				: https://leetcode.com/problems/median-of-two-sorted-arrays
 * Description		: Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the
 * 					  two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * Examples			: Example 1:
 * 					  Input: nums1 = [1,3], nums2 = [2]
 * 					  Output: 2.00000
 * 					  Explanation: merged array = [1,2,3] and median is 2.
 *
 * 					  Example 2:
 * 					  Input: nums1 = [1,2], nums2 = [3,4]
 * 					  Output: 2.50000
 * 					  Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	len1, len2 := len(nums1), len(nums2)
	half := (len1 + len2) / 2
	left, right := 0, len1-1

	// Condition where 'left' and 'right' in the range of
	// 	- the possible minimum part taken from 'nums1' represents by "right = -1", means no values taken from 'nums1'
	// 	- the possible maximum part taken from 'nums1' represents by "left = len of nums1", means all values taken
	// 		from nums1
	for left <= len1 && right >= -1 {
		part1 := (left + right) / 2
		if left+right < 0 {
			part1 = -1
		}

		// Find corresponding partition in nums2. We need (half) elements total on the left side.
		// part1+1 elements come from nums1, so we need (half-1-(part1+1)) from nums2
		part2 := half - part1 - 2

		maxLeft1 := math.Inf(-1) // default value in lower bound (index = -1)
		if part1 >= 0 {
			maxLeft1 = float64(nums1[part1])
		}

		minRight1 := math.Inf(1) // default value in upper bound (index = len of nums1)
		if part1+1 < len1 {
			minRight1 = float64(nums1[part1+1])
		}

		maxLeft2 := math.Inf(-1) // default value in lower bound (index = -1)
		if part2 >= 0 {
			maxLeft2 = float64(nums2[part2])
		}

		minRight2 := math.Inf(1) // default value in upper bound (index = len of nums2)
		if part2+1 < len2 {
			minRight2 = float64(nums2[part2+1])
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (len1+len2)%2 == 1 {
				return min(minRight1, minRight2)
			}
			return (max(maxLeft1, maxLeft2) + min(minRight1, minRight2)) / 2
		} else if maxLeft1 > minRight2 {
			right = part1 - 1
		} else {
			left = part1 + 1
		}
	}

	return 0
}

func RunTestFindMedianSortedArrays() {
	runner.InitMetrics("MedianOfTwoSortedArrays")

	testCases := map[string]struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		"case-1": {
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2.0,
		},
		"case-2": {
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: 2.5,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := runner.ExecCountMetrics(findMedianSortedArrays, tc.nums1, tc.nums2).(float64)
		format.PrintInput(map[string]interface{}{"nums1": tc.nums1, "nums2": tc.nums2})

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
