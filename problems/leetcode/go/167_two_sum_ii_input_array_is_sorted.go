package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Two Sum II - Input Array is Sorted
 * Topics           : Array, Two Pointers, Binary Search
 * Level            : Medium
 * URL              : https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
 * Description      : Given a 1-indexed array sorted in non-decreasing order, you must find the two numbers that add
 * 						up to a given target. Each input guarantees exactly one solution, and you cannot reuse the
 * 						same element twice. Return the pair of indices as 1-based positions.
 * Examples         :
 * 					Example 1:
 * 					Input: numbers = [2,7,11,15], target = 9
 * 					Output: [1,2]
 * 					Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
 *
 * 					Example 2:
 * 					Input: numbers = [2,3,4], target = 6
 * 					Output: [1,3]
 * 					Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
 *
 * 					Example 3:
 * 					Input: numbers = [-1,0], target = -1
 * 					Output: [1,2]
 * 					Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
 */

func twoSumII(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}

	start, end := 0, len(numbers)-1
	for start < end {
		totalSum := numbers[start] + numbers[end]
		if totalSum == target {
			return []int{start + 1, end + 1}
		} else if totalSum > target {
			end -= 1
		} else {
			start += 1
		}
	}

	return []int{}
}

func RunTestTwoSumII() {
	testCases := map[string]struct {
		numbers []int
		target  int
		expect  []int
	}{
		"case-1": {
			numbers: []int{2, 7, 11, 15},
			target:  9,
			expect:  []int{1, 2},
		},
		"case-2": {
			numbers: []int{2, 3, 4},
			target:  6,
			expect:  []int{1, 3},
		},
		"case-3": {
			numbers: []int{-1, 0},
			target:  -1,
			expect:  []int{1, 2},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := twoSumII(testCase.numbers, testCase.target)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
