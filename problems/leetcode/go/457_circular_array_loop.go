package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("CircularArrayLoop", RunTestCircularArrayLoop)
}

/*
 * Problem 			: Circular Array Loop
 * Topics           : Array, Hash Table, Two Pointers
 * Level            : Medium
 * URL              : https://leetcode.com/problems/circular-array-loop
 * Description      : You have a circular array of non-zero integers where each element tells you how many positions
 * 					to jump from that index. Positive values mean jump forward, negative values mean jump backward.
 * 					Since the array is circular, jumping past the end brings you to the beginning, and jumping before
 * 					the start brings you to the end. Determine if there exists a valid cycle in this array. A valid
 * 					cycle must meet three conditions:
 * 						- It forms a loop: Starting from some index, following the jump instructions repeatedly must
 * 							eventually return you to the starting index
 * 						- Consistent direction: All jumps in the cycle must be in the same direction (either all
 * 							forward with positive values or all backward with negative values)
 * 						- Multiple elements: The cycle must contain more than one element (you can't have a cycle that
 * 							just loops on itself)
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [2,-1,1,2,2]
 * 					Output: true
 * 					Explanation: The graph shows how the indices are connected. White nodes are jumping forward, while
 * 					red is jumping backward.
 * 					We can see the cycle 0 --> 2 --> 3 --> 0 --> ..., and all of its nodes are white (jumping in the
 * 					same direction).
 *
 * 					Example 2:
 * 					Input: nums = [-1,-2,-3,-4,-5,6]
 * 					Output: false
 * 					Explanation: The graph shows how the indices are connected. White nodes are jumping forward, while
 * 					red is jumping backward.
 * 					The only cycle is of size 1, so we return false.
 *
 * 					Example 3:
 * 					Input: nums = [1,-1,5,1,4]
 * 					Output: true
 * 					Explanation: The graph shows how the indices are connected. White nodes are jumping forward, while
 * 					red is jumping backward.
 * 					We can see the cycle 0 --> 1 --> 0 --> ..., and while it is of size > 1, it has a node jumping
 * 					forward and a node jumping backward, so it is not a cycle.
 * 					We can see the cycle 3 --> 4 --> 3 --> ..., and all of its nodes are white (jumping in the same
 * 					direction).
 * Reference		: https://algo.monster/liteproblems/457
 */

func circularArrayLoop(nums []int) bool {
	slowIdx, fastIdx := 0, 0
	numsLen := len(nums)

	getNextIndex := func(i int) int {
		return ((i+nums[i])%numsLen + numsLen) % numsLen
	}

	for idx := range nums {
		slowIdx = idx
		fastIdx = getNextIndex(idx)

		// check whether current number and next two numbers have the same direction
		for (nums[slowIdx]*nums[fastIdx] > 0) &&
			(nums[slowIdx]*nums[getNextIndex(fastIdx)] > 0) {

			if slowIdx == fastIdx {
				if slowIdx != getNextIndex(slowIdx) {
					return true
				}
				break
			}

			fastIdx = getNextIndex(getNextIndex(fastIdx))
			slowIdx = getNextIndex(slowIdx)
		}
	}

	return false
}

func RunTestCircularArrayLoop() {
	testCases := map[string]struct {
		nums   []int
		expect bool
	}{
		"case-1": {
			nums:   []int{2, -1, 1, 2, 2},
			expect: true,
		},
		"case-2": {
			nums:   []int{-1, -2, -3, -4, -5, 6},
			expect: false,
		},
		"case-3": {
			nums:   []int{1, -1, 5, 1, 4},
			expect: true,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := circularArrayLoop(testCase.nums)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
