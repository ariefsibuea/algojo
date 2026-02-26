package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ReorderList", RunTestReorderList)
}

/*
 * Problem	: Reorder List
 * Topics	: Linked List, Two Pointers, Stack, Recursion
 * Level	: Medium
 * URL		: https://leetcode.com/problems/reorder-list
 *
 * Description:
 * 		You are given the head of a singly linked-list. The list can be represented as:
 *
 * 		L0 → L1 → … → Ln - 1 → Ln
 *
 * 		Reorder the list to be on the following form:
 *
 * 		L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 		You may not modify the values in the list's nodes. Only nodes themselves may be changed.
 *
 * Constraints:
 * 		- The number of nodes in the list is in the range [1, 10^5]
 * 		- 1 <= Node.val <= 1000
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [1,2,3,4]
 * 		Output: [1,4,2,3]
 *
 * 		Example 2:
 * 		Input: head = [1,2,3,4,5]
 * 		Output: [1,5,2,4,3]
 */

func reorderList(head *ListNode) {
	reorderListSolutions.withExtraSpace(head)
}

type reorderListSolution struct{}

var reorderListSolutions = reorderListSolution{}

func (s *reorderListSolution) withTwoPointers(head *ListNode) {
	// find the middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse node direction from middle -> end
	prev, curr := (*ListNode)(nil), slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// tailor the nodes' direction
	firstPtr, secondPtr := head, prev
	for firstPtr != nil && secondPtr != nil {
		next := firstPtr.Next
		firstPtr.Next = secondPtr
		firstPtr = next

		next = secondPtr.Next
		secondPtr.Next = firstPtr
		secondPtr = next
	}
}

func (s *reorderListSolution) withExtraSpace(head *ListNode) {
	// store nodes into slice
	nodes := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	// tailor the nodes' direction
	start, end := 0, len(nodes)-1
	for start < end {
		nodes[start].Next = nodes[end]
		start++

		if start < end {
			nodes[end].Next = nodes[start]
			end--
		}
	}
	nodes[start].Next = nil
}

func RunTestReorderList() {
	runner.InitMetrics("ReorderList")

	testCases := map[string]struct {
		head   *ListNode
		expect []int
	}{
		"four-nodes-even": {
			head:   buildReorderListFourNodesHead(),
			expect: []int{1, 4, 2, 3},
		},
		"five-nodes-odd": {
			head:   buildReorderListFiveNodesHead(),
			expect: []int{1, 5, 2, 4, 3},
		},
		"single-node": {
			head:   buildReorderListSingleNodeHead(),
			expect: []int{1},
		},
		"two-nodes": {
			head:   buildReorderListTwoNodesHead(),
			expect: []int{1, 2},
		},
		"six-nodes-even": {
			head:   buildReorderListSixNodesHead(),
			expect: []int{1, 6, 2, 5, 3, 4},
		},
	}

	var passedCount int
	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		headSlice := func(h *ListNode) []int {
			var s []int
			for ; h != nil; h = h.Next {
				s = append(s, h.Val)
			}
			return s
		}
		format.PrintInput(map[string]interface{}{"head": headSlice(testCase.head)})

		runner.ExecCountMetrics(reorderList, testCase.head)

		arrResult := make([]int, 0)
		current := testCase.head
		for current != nil {
			arrResult = append(arrResult, current.Val)
			current = current.Next
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

func buildReorderListFourNodesHead() *ListNode {
	node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	return &ListNode{Val: 1, Next: node2}
}

func buildReorderListFiveNodesHead() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	return &ListNode{Val: 1, Next: node2}
}

func buildReorderListSingleNodeHead() *ListNode {
	return &ListNode{Val: 1}
}

func buildReorderListTwoNodesHead() *ListNode {
	node2 := &ListNode{Val: 2}
	return &ListNode{Val: 1, Next: node2}
}

func buildReorderListSixNodesHead() *ListNode {
	node6 := &ListNode{Val: 6}
	node5 := &ListNode{Val: 5, Next: node6}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	return &ListNode{Val: 1, Next: node2}
}
