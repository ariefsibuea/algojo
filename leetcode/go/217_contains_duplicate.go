package main

import (
	"fmt"
	"os"
)

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
		if !EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
