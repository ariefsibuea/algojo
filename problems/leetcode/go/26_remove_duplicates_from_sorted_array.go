package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("RemoveDuplicatesFromSortedArray", RunTestRemoveDuplicatesFromSortedArray)
}

/*
 * Problem	: Remove Duplicates from Sorted Array
 * Topics	: Array, Two Pointers
 * Level	: Easy
 * URL		: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
 *
 * Description:
 * 		Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such
 * 		that each unique element appears only once. The relative order of the elements should be kept the
 * 		same. Then return the number of unique elements in nums.
 *
 * Constraints:
 * 		- 1 <= nums.length <= 3 * 10^4
 * 		- -100 <= nums[i] <= 100
 * 		- nums is sorted in non-decreasing order.
 *
 * Examples:
 * 		Example 1:
 * 		Input: nums = [1,1,2]
 * 		Output: 2, nums = [1,2,_]
 * 		Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2
 * 		respectively.
 *
 * 		Example 2:
 * 		Input: nums = [0,0,1,1,1,2,2,3,3,4]
 * 		Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
 * 		Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1,
 * 		2, 3, and 4 respectively.
 */

func removeDuplicates(nums []int) int {
	currentIdx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[currentIdx] {
			nums[currentIdx+1] = nums[i]
			currentIdx = currentIdx + 1
		}
	}

	return currentIdx + 1
}

func RunTestRemoveDuplicatesFromSortedArray() {
	runner.InitMetrics("RemoveDuplicatesFromSortedArray")

	testCases := map[string]struct {
		nums       []int
		expectNums []int
		expect     int
	}{
		"example-1-basic": {
			nums:       []int{1, 1, 2},
			expectNums: []int{1, 2},
			expect:     2,
		},
		"example-2-multiple-duplicates": {
			nums:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectNums: []int{0, 1, 2, 3, 4},
			expect:     5,
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"nums": tc.nums})

		nums := make([]int, len(tc.nums))
		copy(nums, tc.nums)

		result := runner.ExecCountMetrics(removeDuplicates, nums).(int)

		if !cmp.EqualNumbers(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		failed := false
		for i, n := range tc.expectNums {
			if !cmp.EqualNumbers(n, nums[i]) {
				format.PrintFailed("array check failed at index-%d: expect = %v - got = %v", i, n, nums[i])
				failed = true
				break
			}
		}
		if failed {
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
