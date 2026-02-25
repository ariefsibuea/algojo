package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("ReorderList", RunTestReorderList)
}

/*
 * Problem 			: Reorder List
 * Topics           : Linked List, Two Pointers, Stack, Recursion
 * Level            : Medium
 * URL              : https://leetcode.com/problems/reorder-list
 * Description      : You are given the head of a singly linked-list. The list can be represented as:
 *
 *                    L0 → L1 → … → Ln - 1 → Ln
 *
 *                    Reorder the list to be on the following form:
 *
 *                    L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 *                    You may not modify the values in the list's nodes. Only nodes themselves may be changed.
 * Examples         :
 * 					  Example 1:
 * 					  Input: head = [1,2,3,4]
 * 					  Output: [1,4,2,3]
 *
 * 					  Example 2:
 * 					  Input: head = [1,2,3,4,5]
 * 					  Output: [1,5,2,4,3]
 */

func reorderList(head *ListNode) {
	var slow, fast = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var secondPtr = slow.Next
	slow.Next = nil
	var previous, temp *ListNode

	for secondPtr != nil {
		temp = secondPtr.Next
		secondPtr.Next = previous
		previous = secondPtr
		secondPtr = temp
	}
	secondPtr = previous

	var firstPtr = head

	for secondPtr != nil {
		temp = firstPtr.Next
		firstPtr.Next = secondPtr
		firstPtr = temp

		temp = secondPtr.Next
		secondPtr.Next = firstPtr
		secondPtr = temp
	}
}

func RunTestReorderList() {
	testCases := map[string]struct {
		head   *ListNode
		expect []int
	}{
		"case-1": {
			head:   mockInputP143Case1(),
			expect: []int{1, 4, 2, 3},
		},
		"case-2": {
			head:   mockInputP143Case2(),
			expect: []int{1, 5, 2, 4, 3},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		reorderList(testCase.head)

		arrResult := make([]int, 0)
		current := testCase.head
		for current != nil {
			arrResult = append(arrResult, current.Val)
			current = current.Next
		}

		if !cmp.EqualSlices(arrResult, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, arrResult)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP143Case1() *ListNode {
	node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	return &ListNode{Val: 1, Next: node2}
}

func mockInputP143Case2() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	return &ListNode{Val: 1, Next: node2}
}
