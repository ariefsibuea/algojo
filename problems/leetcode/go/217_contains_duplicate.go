package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Contains Duplicate
 * Topics           : Array, Hash Table, Sorting
 * Level            : Easy
 * URL              : https://leetcode.com/problems/contains-duplicate
 * Description      : Given an integer array nums, return true if any value appears at least twice in the array, and
 * 					return false if every element is distinct.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [1,2,3,1]
 * 					Output: true
 *
 * 					Example 2:
 * 					Input: nums = [1,2,3,4]
 * 					Output: false
 *
 * 					Example 3:
 * 					Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * 					Output: true
 */

func containsDuplicate(nums []int) bool {
	hasExist := make(map[int]bool)
	for _, num := range nums {
		if hasExist[num] {
			return true
		}
		hasExist[num] = true
	}
	return false
}

func RunTestContainsDuplicate() {
	testCases := map[string]struct {
		nums   []int
		expect bool
	}{
		"case-1": {
			nums:   []int{1, 2, 3, 1},
			expect: true,
		},
		"case-2": {
			nums:   []int{1, 2, 3, 4},
			expect: false,
		},
		"case-3": {
			nums:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expect: true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := containsDuplicate(testCase.nums)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
