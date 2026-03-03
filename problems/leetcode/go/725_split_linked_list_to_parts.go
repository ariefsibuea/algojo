package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SplitLinkedListToParts", RunTestSplitLinkedListToParts)
}

/*
 * Problem	: Split Linked List in Parts
 * Topics	: Linked List
 * Level	: Medium
 * URL		: https://leetcode.com/problems/split-linked-list-in-parts/
 *
 * Description:
 * 		Given the head of a singly linked list and an integer k, split the linked list into k consecutive
 * 		parts. The length of each part should be as equal as possible: no two parts should have a size
 * 		differing by more than one. This may lead to some parts being null. The parts should be in the
 * 		order of occurrence in the input list, and parts occurring earlier should always have a size
 * 		greater than or equal to parts occurring later. Return an array of k parts.
 *
 * Constraints:
 * 		- 1 <= k <= 50
 * 		- 0 <= Node.val <= 1000
 * 		- The number of nodes in the list is in the range [0, 1000]
 *
 * Examples:
 * 		Example 1:
 * 		Input: root = [1, 2, 3], k = 5
 * 		Output: [[1],[2],[3],[],[]]
 * 		Explanation: With only 3 nodes but needing 5 parts, the first three parts get one node each
 * 		and the remaining two parts are null (empty).
 *
 * 		Example 2:
 * 		Input: root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
 * 		Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
 * 		Explanation: With 10 nodes split into 3 parts, the first part gets 4 nodes (10/3 = 3 with
 * 		remainder 1, so first 1 part gets +1), and the remaining two parts get 3 nodes each.
 */

func splitListToParts(head *ListNode, k int) []*ListNode {
	parts := make([]*ListNode, k)
	if head == nil {
		return parts
	}

	num := 0
	for curr := head; curr != nil; curr = curr.Next {
		num++
	}

	partSize := num / k
	remainder := num % k

	curr := head
	for i := 0; i < k && curr != nil; i++ {
		parts[i] = curr

		size := partSize
		if remainder > 0 {
			size++
			remainder--
		}

		for j := 1; j < size; j++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = nil
		curr = next
	}

	return parts
}

func RunTestSplitLinkedListToParts() {
	runner.InitMetrics("SplitLinkedListToParts")

	testCases := map[string]struct {
		head   *ListNode
		k      int
		expect [][]int
	}{
		"example-1-more-parts-than-nodes": {
			head: NewListFromSlice([]int{1, 2, 3}),
			k:    5,
			expect: [][]int{
				{1},
				{2},
				{3},
				nil,
				nil,
			},
		},
		"example-2-balanced-split": {
			head: NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			k:    3,
			expect: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7},
				{8, 9, 10},
			},
		},
		"single-node-k-equals-one": {
			head: NewListFromSlice([]int{1}),
			k:    1,
			expect: [][]int{
				{1},
			},
		},
		"k-equals-number-of-nodes": {
			head: NewListFromSlice([]int{1, 2, 3, 4}),
			k:    4,
			expect: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		"empty-list": {
			head: nil,
			k:    3,
			expect: [][]int{
				nil,
				nil,
				nil,
			},
		},
		"k-greater-than-nodes-with-remainder": {
			head: NewListFromSlice([]int{1, 2, 3, 4, 5}),
			k:    8,
			expect: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
				nil,
				nil,
				nil,
			},
		},
		"two-nodes-three-parts": {
			head: NewListFromSlice([]int{1, 2}),
			k:    3,
			expect: [][]int{
				{1},
				{2},
				nil,
			},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"head": testCase.head, "k": testCase.k})

		result := runner.ExecCountMetrics(splitListToParts, testCase.head, testCase.k).([]*ListNode)

		arrResult := make([][]int, len(result))
		for i, node := range result {
			arrResult[i] = listNodeToSlice(node)
		}

		if !cmp.EqualSlices(arrResult, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, arrResult)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
