package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("RemoveNthNodeFromEndOfList", RunTestRemoveNthNodeFromEndOfList)
}

/*
 * Problem	: Remove Nth Node From End of List
 * Topics	: Linked List, Two Pointers
 * Level	: Medium
 * URL		: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 *
 * Description:
 * 		Given the head of a linked list, remove the nth node from the end of the list and return its
 * 		head.
 *
 * Constraints:
 * 		- The number of nodes in the list is sz.
 * 		- 1 <= sz <= 30
 * 		- 0 <= Node.val <= 100
 * 		- 1 <= n <= sz
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [1,2,3,4,5], n = 2
 * 		Output: [1,2,3,5]
 *
 * 		Example 2:
 * 		Input: head = [1], n = 1
 * 		Output: []
 *
 * 		Example 3:
 * 		Input: head = [1,2], n = 1
 * 		Output: [1]
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{
		Val:  0,
		Next: head,
	}

	fast := preHead
	slow := preHead

	i := 0
	for i < n && fast.Next != nil {
		fast = fast.Next
		i += 1
	}
	if i < n-1 {
		return preHead.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return preHead.Next
}

func RunTestRemoveNthNodeFromEndOfList() {
	runner.InitMetrics("RemoveNthNodeFromEndOfList")

	linkedList1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	linkedList2 := ListNode{Val: 1}

	testCases := map[string]struct {
		head   *ListNode
		n      int
		expect []int
	}{
		"example-1-basic": {
			head:   &linkedList1,
			n:      2,
			expect: []int{1, 2, 3, 5},
		},
		"example-2-single-node": {
			head:   &linkedList2,
			n:      1,
			expect: []int{},
		},
	}

	var passedCount int
	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"head": tc.head, "n": tc.n})

		head := runner.ExecCountMetrics(removeNthFromEnd, tc.head, tc.n).(*ListNode)
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
