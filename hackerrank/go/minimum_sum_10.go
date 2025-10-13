package main

import (
	"container/heap"
	"fmt"
	"os"
)

/**
 * Problem 			: Minimum Sum 10
 * Level            : Medium
 * URL              : https://www.hackerrank.com/challenges/minimum-sum-11
 * Description      :
 * Examples         :
 */

type MaxHeap32 []int32

func (h MaxHeap32) Len() int { return len(h) }

// Less flips the comparison so that the largest item is “less” in this interface => comes first
func (h MaxHeap32) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap32) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap32) Push(x interface{}) {
	*h = append(*h, x.(int32))
}

func (h *MaxHeap32) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minSum10(num []int32, k int32) int32 {
	h := MaxHeap32(num)
	heap.Init(&h)

	for i := 0; i < int(k); i++ {
		maxVal := heap.Pop(&h).(int32)
		newVal := (maxVal + 1) / 2
		heap.Push(&h, newVal)
	}

	sum := int32(0)
	for _, v := range h {
		sum += v
	}

	return int32(sum)
}

func RunTestMinimumSum10() {
	testCases := map[string]struct {
		nums   []int32
		k      int32
		expect int32
	}{
		"case-1": {
			nums:   []int32{2},
			k:      1,
			expect: 1,
		},
		"case-2": {
			nums:   []int32{2, 3},
			k:      1,
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := minSum10(testCase.nums, testCase.k)
		if !EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
