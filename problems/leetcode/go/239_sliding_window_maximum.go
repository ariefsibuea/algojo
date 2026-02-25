package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("SlidingWindowMaximum", RunTestSlidingWindowMaximum)
}

/*
 * Problem 			: Sliding Window Maximum
 * Topics           : Array, Queue, Sliding Window, Heap (Priority Queue), Monotonic Queue, Deque
 * Level            : Hard
 * URL              : https://leetcode.com/problems/sliding-window-maximum
 * Description      : You are given an array of integers `nums`, there is a sliding window of size `k` which is moving
 * 					  from the very left of the array to the very right. You can only see the k numbers in the window.
 *                    Each time the sliding window moves right by one position.
 *                    Return an array of the maximum value for each window.
 * Constraints      :
 *                    - 1 <= nums.length <= 10^5
 *                    - -10^4 <= nums[i] <= 10^4
 *                    - 1 <= k <= nums.length
 * Examples         :
 *                    Example 1:
 *                    Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
 *                    Output: [3,3,5,5,6,7]
 *                    Explanation:
 *                    Window position                Max
 *                    ---------------               -----
 *                    [1  3  -1] -3  5  3  6  7       3
 *                     1 [3  -1  -3] 5  3  6  7       3
 *                     1  3 [-1  -3  5] 3  6  7       5
 *                     1  3  -1 [-3  5  3] 6  7       5
 *                     1  3  -1  -3 [5  3  6] 7       6
 *                     1  3  -1  -3  5 [3  6  7]      7
 *
 *                    Example 2:
 *                    Input: nums = [1], k = 1
 *                    Output: [1]
 */

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	result := make([]int, 0, n-k+1)
	deque := make([]int, 0)
	dequeLen := 0

	for i := 0; i < n; i++ {
		dequeLen = len(deque)

		for dequeLen > 0 && nums[deque[dequeLen-1]] <= nums[i] {
			deque = deque[:dequeLen-1]
			dequeLen = len(deque)
		}

		deque = append(deque, i)

		if deque[0] <= i-k {
			deque = deque[1:]
		}

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func RunTestSlidingWindowMaximum() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect []int
	}{
		"example-from-problem": {
			nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:      3,
			expect: []int{3, 3, 5, 5, 6, 7},
		},
		"single-element-array": {
			nums:   []int{1},
			k:      1,
			expect: []int{1},
		},
		"k-is-equal-to-array-length": {
			nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:      8,
			expect: []int{7},
		},
		"k-is-one": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      1,
			expect: []int{1, 2, 3, 4, 5},
		},
		"descending-order": {
			nums:   []int{5, 4, 3, 2, 1},
			k:      3,
			expect: []int{5, 4, 3},
		},
		"ascending-order": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      3,
			expect: []int{3, 4, 5},
		},
		"all-elements-are-the-same": {
			nums:   []int{5, 5, 5, 5, 5},
			k:      2,
			expect: []int{5, 5, 5, 5},
		},
		"contains-negative-numbers": {
			nums:   []int{-1, -2, -3, -4, -5},
			k:      2,
			expect: []int{-1, -2, -3, -4},
		},
		"mixed-positive-and-negative-numbers": {
			nums:   []int{1, -1, 2, -2, 3, -3},
			k:      3,
			expect: []int{2, 2, 3, 3},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := maxSlidingWindow(testCase.nums, testCase.k)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
