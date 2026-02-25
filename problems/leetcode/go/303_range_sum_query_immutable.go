package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("RangeSumQueryImmutable", RunTestRangeSumQueryImmutable)
}

/*
 * Problem 			: Range Sum Query - Immutable
 * Topics           : Array, Design, Prefix Sum
 * Level            : Easy
 * URL              : https://leetcode.com/problems/range-sum-query-immutable
 * Description      : <Description>
 * Examples         : <Examples>
 */

type NumArray struct {
	prefix []int
}

func NumArrayConstructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefix[i+1] = sum
	}

	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

func RunTestRangeSumQueryImmutable() {
	testCases := map[string]struct{}{}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		fmt.Println(testCase)
		// result := twoSum(testCase.nums, testCase.target)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	os.Exit(1)
		// }
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
