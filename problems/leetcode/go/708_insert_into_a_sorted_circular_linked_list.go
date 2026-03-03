package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("InsertIntoASortedCircularLinkedList", RunTestInsertIntoASortedCircularLinkedList)
}

/*
 * Problem	: Insert into a Sorted Circular Linked List
 * Topics	: Linked List
 * Level	: Medium
 * URL		: https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/
 *
 * Description:
 * 		Given a circular linked list sorted in non-descending order, insert a new value into the list while maintaining
 * 		the sorted order. The last node of the list points back to the first node, forming a circle. You are given a
 * 		reference to any node in the list. If the list is empty (head is null), create a new single-node circular list
 * 		with the insert value and return it. Otherwise, insert the value at the appropriate position and return the
 * 		original head node.
 *
 * Constraints:
 * 		- The list is sorted in non-descending order (ascending with possible duplicates)
 * 		- The list is circular (last node.next points to first node, forming a loop)
 * 		- You receive a reference to any node in the list, not necessarily the head
 * 		- If the list is empty, create a new single-node circular list
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [3,4,5], insertVal = 1 (starting from node 4)
 * 		Output: [3,4,5,1] or [1,3,4,5]
 * 		Explanation: The list is 4 -> 5 -> 3 -> (back to 4). Inserting 1 at the end maintains
 * 		sorted order since 1 < 3 and 1 > 5.
 *
 * 		Example 2:
 * 		Input: head = [], insertVal = 1
 * 		Output: [1]
 * 		Explanation: Empty list case - create a new circular list with single node.
 *
 * 		Example 3:
 * 		Input: head = [1,1,1], insertVal = 2
 * 		Output: [1,1,1,2]
 * 		Explanation: Insert 2 after the last 1 to maintain sorted order.
 */

func insertIntoSortedCircularLinkedList(head *ListNode, insertVal int) *ListNode {
	if head == nil {
		newNode := &ListNode{Val: insertVal, Next: nil}
		newNode.Next = newNode
		return newNode
	}

	prev, next := head, head.Next
	for next != head {
		if prev.Val <= insertVal && insertVal <= next.Val {
			break
		}
		if prev.Val > next.Val && (prev.Val < insertVal || insertVal < next.Val) {
			break
		}

		prev = next
		next = next.Next
	}

	newNode := &ListNode{Val: insertVal, Next: next}
	prev.Next = newNode

	return head
}

func RunTestInsertIntoASortedCircularLinkedList() {
	runner.InitMetrics("InsertIntoASortedCircularLinkedList")

	testCases := map[string]struct {
		head      *ListNode
		insertVal int
		expect    []int
	}{
		"example-1-insert-min-value": {
			head:      buildInsertIntoSortedCircularLinkedListExample1Head(),
			insertVal: 1,
			expect:    []int{4, 5, 1, 3},
		},
		"example-2-empty-list": {
			head:      nil,
			insertVal: 1,
			expect:    []int{1},
		},
		"example-3-insert-max-value": {
			head:      buildInsertIntoSortedCircularLinkedListExample3Head(),
			insertVal: 2,
			expect:    []int{1, 1, 1, 2},
		},
		"insert-middle-ascending": {
			head:      buildInsertIntoSortedCircularLinkedListInsertMiddleHead(),
			insertVal: 3,
			expect:    []int{1, 2, 3, 4, 5},
		},
		"insert-duplicate-value": {
			head:      buildInsertIntoSortedCircularLinkedListInsertDuplicateHead(),
			insertVal: 2,
			expect:    []int{1, 2, 2, 3},
		},
		"edge-case-single-node": {
			head:      buildInsertIntoSortedCircularLinkedListSingleNodeHead(),
			insertVal: 0,
			expect:    []int{1, 0},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"head": circularListToSlice(testCase.head), "insertVal": testCase.insertVal})

		result := runner.ExecCountMetrics(insertIntoSortedCircularLinkedList, testCase.head, testCase.insertVal).(*ListNode)

		actual := circularListToSlice(result)
		if !cmp.EqualSlices(actual, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, actual)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

// Mock input functions for test cases

func buildInsertIntoSortedCircularLinkedListExample1Head() *ListNode {
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node3.Next = node4
	node4.Next = node5
	node5.Next = node3
	return node4 // Start from node 4 as per example
}

func buildInsertIntoSortedCircularLinkedListExample3Head() *ListNode {
	node1a := &ListNode{Val: 1}
	node1b := &ListNode{Val: 1}
	node1c := &ListNode{Val: 1}
	node1a.Next = node1b
	node1b.Next = node1c
	node1c.Next = node1a
	return node1a
}

func buildInsertIntoSortedCircularLinkedListInsertMiddleHead() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node4
	node4.Next = node5
	node5.Next = node1
	return node1
}

func buildInsertIntoSortedCircularLinkedListInsertDuplicateHead() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1
	return node1
}

func buildInsertIntoSortedCircularLinkedListSingleNodeHead() *ListNode {
	node1 := &ListNode{Val: 1}
	node1.Next = node1
	return node1
}
