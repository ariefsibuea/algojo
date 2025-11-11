package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * LeetCode Problem : Remove Duplicates from Sorted Array
 * Topics           : Array, Two Pointers
 * Level            : Easy
 * URL              : https://leetcode.com/problems/remove-duplicates-from-sorted-array
 * Description      : Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such
 * 					that each unique element appears only once. The relative order of the elements should be kept the
 * 					same. Then return the number of unique elements in nums.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [1,1,2]
 * 					Output: 2, nums = [1,2,_]
 * 					Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2
 * 					respectively.
 *
 * 					Example 2:
 * 					Input: nums = [0,0,1,1,1,2,2,3,3,4]
 * 					Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
 * 					Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1,
 * 					2, 3, and 4 respectively.
 */

func removeDuplicates(nums []int) int {
	currentIdx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[currentIdx] {
			nums[currentIdx+1] = nums[i]
			currentIdx = currentIdx + 1
		}
	}

	return currentIdx + 1
}

func RunTestRemoveDuplicatesFromSortedArray() {
	testCases := map[string]struct {
		nums       []int
		expectNums []int
		expect     int
	}{
		"case-1": {
			nums:       []int{1, 1, 2},
			expectNums: []int{1, 2},
			expect:     2,
		},
		"case-2": {
			nums:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectNums: []int{0, 1, 2, 3, 4},
			expect:     5,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := removeDuplicates(testCase.nums)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		for i, n := range testCase.expectNums {
			if !cmp.EqualNumbers(n, testCase.nums[i]) {
				fmt.Printf("==== FAILED: expect = %v - got = %v in the index-%d\n", testCase.expect, result, i)
				os.Exit(1)
			}
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
