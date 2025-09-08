package main

import (
	"fmt"
	"os"
)

func topKFrequent(nums []int, k int) []int {
	frequent := make(map[int]int, len(nums))
	for _, num := range nums {
		frequent[num]++
	}

	maxFrequent := 0
	bucket := make([][]int, len(nums)+1)
	for num, freq := range frequent {
		maxFrequent = max(maxFrequent, freq)
		bucket[freq] = append(bucket[freq], num)
	}

	result := make([]int, 0)
	for i := maxFrequent; i >= 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		result = append(result, bucket[i]...)
		if len(result) >= k {
			break
		}
	}

	return result[:k]
}

func RunTestTopKFrequentElements() {
	testCases := map[string]struct {
		nums   []int
		k      int
		expect []int
	}{
		"case-1": {
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			expect: []int{1, 2},
		},
		"case-2": {
			nums:   []int{1},
			k:      1,
			expect: []int{1},
		},
		"case-3": {
			nums:   []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:      2,
			expect: []int{1, 2},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := topKFrequent(testCase.nums, testCase.k)
		if !EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
