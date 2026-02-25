package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("MiddleOfTheLinkedList", RunTestMiddleOfTheLinkedList)
}

/*
 * Problem 			: Middle of the Linked List
 * Topics           : Linked List, Two Pointers
 * Level            : Easy
 * URL              : https://leetcode.com/problems/middle-of-the-linked-list
 * Description      : Given the head of a singly linked list, return the middle node of the linked list. If there are
 * 					two middle nodes, return the second middle node.
 * Examples         :
 * 					Example 1:
 * 					Input: head = [1,2,3,4,5]
 * 					Output: [3,4,5]
 * 					Explanation: The middle node of the list is node 3.
 *
 * 					Example 2:
 * 					Input: head = [1,2,3,4,5,6]
 * 					Output: [4,5,6]
 * 					Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
 */

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func RunTestMiddleOfTheLinkedList() {
	testCases := map[string]struct {
		head   *ListNode
		expect int
	}{
		"case-1": {
			head:   mockHeadP876Case1(),
			expect: 3,
		},
		"case-2": {
			head:   mockHeadP876Case2(),
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := middleNode(testCase.head)
		if !cmp.EqualNumbers(result.Val, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result.Val)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockHeadP876Case1() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}

func mockHeadP876Case2() *ListNode {
	node6 := &ListNode{Val: 6}
	node5 := &ListNode{Val: 5, Next: node6}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}
