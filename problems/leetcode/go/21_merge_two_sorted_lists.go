package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("MergeTwoSortedLists", RunTestMergeTwoSortedLists)
}

/*
 * Problem	: Merge Two Sorted Lists
 * Topics	: Linked List, Recursion
 * Level	: Easy
 * URL		: https://leetcode.com/problems/merge-two-sorted-lists/
 *
 * Description:
 * 		You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one
 * 		sorted list. The list should be made by splicing together the nodes of the first two lists. Return
 * 		the head of the merged linked list.
 *
 * Constraints:
 * 		- The number of nodes in both lists is in the range [0, 50].
 * 		- -100 <= Node.val <= 100
 * 		- Both list1 and list2 are sorted in non-decreasing order.
 *
 * Examples:
 * 		Example 1:
 * 		Input: list1 = [1,2,4], list2 = [1,3,4]
 * 		Output: [1,1,2,3,4,4]
 *
 * 		Example 2:
 * 		Input: list1 = [], list2 = []
 * 		Output: []
 *
 * 		Example 3:
 * 		Input: list1 = [], list2 = [0]
 * 		Output: [0]
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func RunTestMergeTwoSortedLists() {
	runner.InitMetrics("MergeTwoSortedLists")

	list1Case1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2Case1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2Case3 := &ListNode{
		Val: 0,
	}

	testCases := map[string]struct {
		list1  *ListNode
		list2  *ListNode
		expect []int
	}{
		"example-1-basic": {
			list1:  list1Case1,
			list2:  list2Case1,
			expect: []int{1, 1, 2, 3, 4, 4},
		},
		"example-2-empty": {
			list1:  nil,
			list2:  nil,
			expect: []int{},
		},
		"example-3-one-empty": {
			list1:  nil,
			list2:  list2Case3,
			expect: []int{0},
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"list1": tc.list1, "list2": tc.list2})

		head := runner.ExecCountMetrics(mergeTwoLists, tc.list1, tc.list2).(*ListNode)
		result := []int{}

		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}

		if !cmp.EqualSlices(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
