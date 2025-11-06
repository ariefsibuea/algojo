package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Squares of a Sorted Array
 * Topics           : Array, Two Pointers, Sorting
 * Level            : Easy
 * URL              : https://leetcode.com/problems/squares-of-a-sorted-array
 * Description      : You are given an integer array 'nums' sorted in non-decreasing order. Your task is to return a
 * 					new array containing the squares of each number, also sorted in non-decreasing order.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [-4,-1,0,3,10]
 * 					Output: [0,1,9,16,100]
 * 					Explanation: After squaring, the array becomes [16,1,0,9,100].
 * 					After sorting, it becomes [0,1,9,16,100].
 *
 * 					Example 2:
 * 					Input: nums = [-7,-3,2,3,11]
 * 					Output: [4,9,9,49,121]
 */

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	leftSqrt, rightSqrt := 0, 0
	left, right := 0, len(nums)-1
	index := right

	for left <= right {
		leftSqrt = nums[left] * nums[left]
		rightSqrt = nums[right] * nums[right]

		if leftSqrt > rightSqrt {
			result[index] = leftSqrt
			left += 1
		} else {
			result[index] = rightSqrt
			right -= 1
		}
		index -= 1
	}

	return result
}

func RunTestSquaresOfASortedArray() {
	testCases := map[string]struct {
		nums   []int
		expect []int
	}{
		"case-1": {
			nums:   []int{-4, -1, 0, 3, 10},
			expect: []int{0, 1, 9, 16, 100},
		},
		"case-2": {
			nums:   []int{-7, -3, 2, 3, 11},
			expect: []int{4, 9, 9, 49, 121},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := sortedSquares(testCase.nums)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
