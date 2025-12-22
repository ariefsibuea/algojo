package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Single Number
 * Topics           : Array, Bit Manipulation
 * Level            : Easy
 * URL              : https://leetcode.com/problems/single-number
 * Description      : <Description>
 * Examples         : <Examples>
 */

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func RunTestSingleNumber() {
	testCases := map[string]struct {
		nums   []int
		expect int
	}{
		"case-1": {
			nums:   []int{2, 2, 1},
			expect: 1,
		},
		"case-2": {
			nums:   []int{4, 1, 2, 1, 2},
			expect: 4,
		},
		"case-3": {
			nums:   []int{1},
			expect: 1,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := singleNumber(testCase.nums)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
