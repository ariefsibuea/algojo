package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Palindrome Linked List
 * Topics           : Linked List, Two Pointers, Stack, Recursion
 * Level            : Easy
 * URL              : https://leetcode.com/problems/palindrome-linked-list
 * Description      : Given the head of a singly linked list. Your task is to determine whether the linked list forms
 * 					a palindrome when reading the values from start to end.
 * Examples         :
 * 					Example 1:
 * 					Input: head = [1,2,2,1]
 * 					Output: true
 *
 * 					Example 2:
 * 					Input: head = [1,2]
 * 					Output: false
 */

type NodeP234 struct {
	Val  int
	Next *NodeP234
}

func isPalindromeLinkedList(head *NodeP234) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev, temp *NodeP234
	var current = slow.Next

	for current != nil {
		temp = current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}

func RunTestPalindromeLinkedList() {
	testCases := map[string]struct {
		head   *NodeP234
		expect bool
	}{
		"case-1": {
			head:   mockHeadP234Case1(),
			expect: true,
		},
		"case-2": {
			head:   mockHeadP234Case2(),
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := isPalindromeLinkedList(testCase.head)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockHeadP234Case1() *NodeP234 {
	node4 := &NodeP234{Val: 1}
	node3 := &NodeP234{Val: 2, Next: node4}
	node2 := &NodeP234{Val: 2, Next: node3}
	head := &NodeP234{Val: 1, Next: node2}
	return head
}

func mockHeadP234Case2() *NodeP234 {
	node2 := &NodeP234{Val: 2}
	head := &NodeP234{Val: 1, Next: node2}
	return head
}
