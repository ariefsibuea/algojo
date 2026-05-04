package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	// register function here
}

/*
 * Problem	: Merge k Sorted Lists
 * Topics	: Linked List, Divide and Conquer, Heap (Priority Queue), Merge Sort
 * Level	: Hard
 * URL		: https://leetcode.com/problems/merge-k-sorted-lists/
 *
 * Description:
 * 		Given an array of k linked-lists where each linked-list is sorted in ascending order, the task is to merge all
 * 		the linked-lists into a single sorted linked-list and return its head. This problem involves combining multiple
 * 		pre-sorted sequences while maintaining the sorted order in the final result.
 *
 * Constraints:
 * 		- k == lists.length
 * 		- 0 <= k <= 10^4
 * 		- 0 <= lists[i].length <= 500
 * 		- -10^4 <= lists[i][j] <= 10^4
 * 		- Each lists[i] is sorted in ascending order
 * 		- The sum of all lists[i].length will not exceed 10^4
 *
 * Examples:
 * 		Example 1:
 * 		Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * 		Output: [1,1,2,3,4,4,5,6]
 * 		Explanation: The three linked-lists are 1->4->5, 1->3->4, and 2->6. Merging them together in sorted order gives
 * 		             1->1->2->3->4->4->5->6.
 *
 * 		Example 2:
 * 		Input: lists = []
 * 		Output: []
 * 		Explanation: When there are no lists to merge, the result is an empty list.
 *
 * 		Example 3:
 * 		Input: lists = [[]]
 * 		Output: []
 * 		Explanation: When the input contains a single empty list, the result is also an empty list.
 */

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}

func RunTestMergeKSortedLists() {
	runner.InitMetrics("MergeKSortedLists")

	testCases := map[string]struct {
		lists  []*ListNode
		expect []int
	}{
		"example-1-three-lists": {
			lists:  buildMergeKSortedListsExample1(),
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		"example-2-empty-lists": {
			lists:  []*ListNode{},
			expect: nil,
		},
		"example-3-single-empty-list": {
			lists:  []*ListNode{nil},
			expect: nil,
		},
		"single-list-only": {
			lists:  buildMergeKSortedListsSingleList(),
			expect: []int{1, 2, 3},
		},
		"all-negative-numbers": {
			lists:  buildMergeKSortedListsNegativeNumbers(),
			expect: []int{-10, -5, -3, -1},
		},
		"lists-with-duplicates": {
			lists:  buildMergeKSortedListsDuplicates(),
			expect: []int{1, 1, 1, 1, 2, 2},
		},
		"multiple-single-element-lists": {
			lists:  buildMergeKSortedListsMultipleSingleElements(),
			expect: []int{1, 2, 3, 4, 5},
		},
		"large-numbers-range": {
			lists:  buildMergeKSortedListsLargeNumbers(),
			expect: []int{-10000, -9999, 0, 9999, 10000},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"lists": listNodeSliceToSlices(testCase.lists)})

		result := runner.ExecCountMetrics(mergeKLists, testCase.lists).(*ListNode)
		resultSlice := listNodeToSlice(result)
		if !cmp.EqualSlices(resultSlice, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, resultSlice)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

// Helper function to convert []*ListNode to [][]int for display
func listNodeSliceToSlices(lists []*ListNode) [][]int {
	result := make([][]int, len(lists))
	for i, list := range lists {
		result[i] = listNodeToSlice(list)
	}
	return result
}

// Mock input functions for Merge k Sorted Lists
func buildMergeKSortedListsExample1() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{1, 4, 5}),
		NewListFromSlice([]int{1, 3, 4}),
		NewListFromSlice([]int{2, 6}),
	}
}

func buildMergeKSortedListsSingleList() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{1, 2, 3}),
	}
}

func buildMergeKSortedListsNegativeNumbers() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{-5, -3}),
		NewListFromSlice([]int{-10, -1}),
	}
}

func buildMergeKSortedListsDuplicates() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{1, 1, 1}),
		NewListFromSlice([]int{1, 2, 2}),
	}
}

func buildMergeKSortedListsMultipleSingleElements() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{1}),
		NewListFromSlice([]int{2}),
		NewListFromSlice([]int{3}),
		NewListFromSlice([]int{4}),
		NewListFromSlice([]int{5}),
	}
}

func buildMergeKSortedListsLargeNumbers() []*ListNode {
	return []*ListNode{
		NewListFromSlice([]int{-10000, 0, 10000}),
		NewListFromSlice([]int{-9999, 9999}),
	}
}
