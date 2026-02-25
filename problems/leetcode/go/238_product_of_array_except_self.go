package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("ProductOfArrayExceptSelf", RunTestProductOfArrayExceptSelf)
}

/*
 * Problem 			: Product of Array Except Self
 * Topics           : Array, Prefix Sum
 * Level            : Medium
 * URL              : https://leetcode.com/problems/product-of-array-except-self
 * Description      : Given an integer array nums, return an array answer such that answer[i] is equal to the product
 * 					  of all the elements of nums except nums[i]. The product of any prefix or suffix of nums is
 * 					  guaranteed to fit in a 32-bit integer. You must write an algorithm that runs in O(n) time and
 * 					  without using the division operation.
 * Examples         :
 * 					  Example 1:
 * 					  Input: nums = [1,2,3,4]
 * 					  Output: [24,12,8,6]
 *
 * 					  Example 2:
 * 					  Input: nums = [-1,1,0,-3,3]
 * 					  Output: [0,0,9,0,0]
 */

func productExceptSelf(nums []int) []int {
	n := len(nums)

	result := make([]int, n)
	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

func RunTestProductOfArrayExceptSelf() {
	testCases := map[string]struct {
		nums   []int
		expect []int
	}{
		"case-1": {
			nums:   []int{1, 2, 3, 4},
			expect: []int{24, 12, 8, 6},
		},
		"case-2": {
			nums:   []int{-1, 1, 0, -3, 3},
			expect: []int{0, 0, 9, 0, 0},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := productExceptSelf(testCase.nums)
		format.PrintInput(map[string]interface{}{"nums": testCase.nums})

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
