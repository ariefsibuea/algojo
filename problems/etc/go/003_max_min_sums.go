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

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = 0
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = 0
	*h = old[0 : n-1]
	return x
}

func maxMinSums(nums []int, n int, k int) []int {
	if k >= n {
		return []int{}
	}

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if k == 0 {
		return []int{totalSum, totalSum}
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, num := range nums {
		if maxHeap.Len() < k {
			heap.Push(maxHeap, num)
		} else if num < (*maxHeap)[0] {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, num)
		}
	}

	sumKSmallest := 0
	for maxHeap.Len() > 0 {
		sumKSmallest += heap.Pop(maxHeap).(int)
	}

	maxSum := totalSum - sumKSmallest

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		if minHeap.Len() < k {
			heap.Push(minHeap, num)
		} else if num > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, num)
		}
	}

	sumKLargest := 0
	for minHeap.Len() > 0 {
		sumKLargest += heap.Pop(minHeap).(int)
	}

	minSum := totalSum - sumKLargest

	return []int{maxSum, minSum}
}

func RunTestMaxMinSums() {
	testCases := map[string]struct {
		nums   []int
		n      int
		k      int
		expect []int
	}{
		"case-1": {
			nums:   []int{1, 2, 3, 4, 5},
			n:      5,
			k:      1,
			expect: []int{14, 10},
		},
		"case-2": {
			nums:   []int{5, 6, 7},
			n:      3,
			k:      2,
			expect: []int{7, 5},
		},
		"case-3": {
			nums:   []int{5, 1, 3, 2, 4},
			n:      5,
			k:      2,
			expect: []int{12, 6},
		},
		"case-4": {
			nums:   []int{10, -1, -2, 5, 1},
			n:      5,
			k:      2,
			expect: []int{16, -2},
		},
		"case-5": {
			nums:   []int{1, 2, 3, 4, 5, 6},
			n:      6,
			k:      3,
			expect: []int{15, 6},
		},
		"case-6": {
			nums:   []int{7, 7, 7, 7},
			n:      4,
			k:      2,
			expect: []int{14, 14},
		},
		"case-7": {
			nums:   []int{0, -5, -10, 20, 30},
			n:      5,
			k:      1,
			expect: []int{45, 5},
		},
		"case-8": {
			nums:   []int{3, 1, 4, 1, 5, 9},
			n:      6,
			k:      3,
			expect: []int{18, 5},
		},
		"case-9": {
			nums:   []int{},
			n:      0,
			k:      3,
			expect: []int{},
		},
		"case-10": {
			nums:   []int{100, -1},
			n:      2,
			k:      2,
			expect: []int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := maxMinSums(testCase.nums, testCase.n, testCase.k)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
