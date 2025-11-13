package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Find the Duplicate Number
 * Topics           : Array, Two Pointers, Binary Search, Bit Manipulation
 * Level            : Medium
 * URL              : https://leetcode.com/problems/find-the-duplicate-number
 * Description      : You are given an array nums containing n + 1 integers, where each integer is in the range [1, n]
 * 					inclusive. The array has exactly one repeated number that appears more than once. Find and return
 * 					this duplicate number.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [1,3,4,2,2]
 * 					Output: 2
 *
 * 					Example 2:
 * 					Input: nums = [3,1,3,4,2]
 * 					Output: 3
 *
 * 					Example 3:
 * 					Input: nums = [3,3,3,3,3]
 * 					Output: 3
 */

func findDuplicate_TortoiseHare(nums []int) int {
	slowIdx, fastIdx := 0, 0
	for {
		fastIdx = nums[nums[fastIdx]]
		slowIdx = nums[slowIdx]
		if slowIdx == fastIdx {
			break
		}
	}

	slowIdx = 0
	for slowIdx != fastIdx {
		fastIdx = nums[fastIdx]
		slowIdx = nums[slowIdx]
	}

	return fastIdx
}

func findDuplicate_MarkVisited(nums []int) int {
	currentIdx := 0
	temp := 0
	result := 0

	for {
		if nums[currentIdx] == 0 {
			result = currentIdx
			break
		}

		temp = nums[currentIdx]
		nums[currentIdx] = 0
		currentIdx = temp
	}

	return result
}

func findDuplicate_HashMap(nums []int) int {
	hasExists := map[int]bool{}

	for _, num := range nums {
		if hasExists[num] {
			return num
		}
		hasExists[num] = true
	}
	return 0
}

func RunTestFindTheDuplicateNumber() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{1, 3, 4, 2, 2},
			expect: 2,
		},
		"case-2": {
			nums:   []int{3, 1, 3, 4, 2},
			expect: 3,
		},
		"case-3": {
			nums:   []int{3, 3, 3, 3, 3},
			expect: 3,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := findDuplicate_TortoiseHare(testCase.nums)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
