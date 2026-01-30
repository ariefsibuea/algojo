package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem          : Split A Circular Linked List
 * Topics           : Linked List
 * Level            : Medium
 * URL              : https://leetcode.com/problems/split-a-circular-linked-list
 * Description      : You are given a circular linked list with an even number of nodes (N). The task is to split this
 *                    list into two equal-sized circular linked lists: one containing the first N/2 nodes and the other
 *                    containing the last N/2 nodes.
 * Constraints      :
 *                    - 1 <= T <= 10
 *                    - 2 <= N <= 10000
 *                    - -10^9 <= node data <= 10^9
 *                    - Time limit: 1 sec
 * Examples         :
 *                    Example 1:
 *                    Input: head = [1, 2, 3, 4] (circular)
 *                    Output: [[1, 2], [3, 4]]
 *                    Explanation: The first N/2 nodes are 1 and 2, which form the first circular list. The last N/2
 *                    nodes are 3 and 4, which form the second circular list.
 */

func splitCircularLinkedList(head *ListNode) []*ListNode {
	slow, fast := head, head.Next

	for fast != head && fast.Next != head {
		fast = fast.Next.Next
		slow = slow.Next
	}

	head2 := slow.Next

	// set first circular
	slow.Next = head

	// set second circular
	fast = head2
	for fast.Next != head {
		fast = fast.Next
	}
	fast.Next = head2

	return []*ListNode{head, head2}
}

func RunTestSplitACircularLinkedList() {
	testCases := map[string]struct {
		head   *ListNode
		expect [][]int
	}{
		"even-nodes-4": {
			head:   inputTestSplitACircularLinkedListCaseEvenNodes4Head(),
			expect: [][]int{{1, 2}, {3, 4}},
		},
		"even-nodes-2": {
			head:   inputTestSplitACircularLinkedListCaseEvenNodes2Head(),
			expect: [][]int{{1}, {2}},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		inputSlice := circularListToSlice(testCase.head)
		result := splitCircularLinkedList(testCase.head)

		format.PrintInput(map[string]interface{}{"head": inputSlice})

		var actual [][]int
		for _, h := range result {
			actual = append(actual, circularListToSlice(h))
		}

		if !cmp.EqualSlices(actual, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, actual)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func circularListToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := []int{head.Val}
	curr := head.Next
	for curr != nil && curr != head {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
}

func inputTestSplitACircularLinkedListCaseEvenNodes4Head() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1
	return node1
}

func inputTestSplitACircularLinkedListCaseEvenNodes2Head() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1
	return node1
}
