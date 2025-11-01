package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * Problem 			: Max and Min Sums
 * Description      : Given an array of integer nums, its length n, and an integer k. Find the maximum and minimum
 * 					possible sums after discarding k numbers. Note that you are not allowed to use a sorting algorithm.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [1, 2, 3, 4, 5], n = 5, k = 1
 * 					Output: [14, 10]
 * 					Explanation: Maximum sum is 14 by discarded 1 and minimum sum is 10 by discarded 5.
 *
 * 					Example 2:
 * 					Input: nums = [5, 6, 7], n = 3, k = 2
 * 					Output:[7, 5]
 * 					Explanation: Maximum sum is 7 by discarded 5 and 6, minimum sum is 5 by discarded 6 and 7.
 */

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = 0
	*h = old[0 : n-1]
	return x
}

func maxMinSums(nums []int, k int) []int {
	if k >= len(nums) {
		return []int{}
	}

	minHeap := &IntMinHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		heap.Push(minHeap, num)
	}

	maxSum, minSum := 0, 0
	num := 0

	for i := 0; i < len(nums); i++ {
		num = heap.Pop(minHeap).(int)
		if i < len(nums)-k {
			minSum += num
		}
		if i >= k {
			maxSum += num
		}
	}

	return []int{maxSum, minSum}
}

func RunTestMaxMinSums() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect []int
	}{
		"case-1": {
			nums:   []int{1, 2, 3, 4, 5},
			k:      1,
			expect: []int{14, 10},
		},
		"case-2": {
			nums:   []int{5, 6, 7},
			k:      2,
			expect: []int{7, 5},
		},
		"case-3": {
			nums:   []int{5, 1, 3, 2, 4},
			k:      2,
			expect: []int{12, 6},
		},
		"case-4": {
			nums:   []int{10, -1, -2, 5, 1},
			k:      2,
			expect: []int{16, -2},
		},
		"case-5": {
			nums:   []int{1, 2, 3, 4, 5, 6},
			k:      3,
			expect: []int{15, 6},
		},
		"case-6": {
			nums:   []int{7, 7, 7, 7},
			k:      2,
			expect: []int{14, 14},
		},
		"case-7": {
			nums:   []int{0, -5, -10, 20, 30},
			k:      1,
			expect: []int{45, 5},
		},
		"case-8": {
			nums:   []int{3, 1, 4, 1, 5, 9},
			k:      3,
			expect: []int{18, 5},
		},
		"case-9": {
			nums:   []int{},
			k:      3,
			expect: []int{},
		},
		"case-10": {
			nums:   []int{100, -1},
			k:      2,
			expect: []int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := maxMinSums(testCase.nums, testCase.k)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
