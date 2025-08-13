package main

import (
	"fmt"
	"os"
)

/*
LeetCode Problem : Remove Nth Node From End of List
Topic            : Linked List, Two Pointers
Level            : Medium
URL              : https://leetcode.com/problems/remove-nth-node-from-end-of-list
Description      : Given the head of a linked list, remove the nth node from the end of the list and return its head.
Examples         :
        Example 1:
        Input: head = [1,2,3,4,5], n = 2
        Output: [1,2,3,5]

        Example 2:
        Input: head = [1], n = 1
        Output: []

        Example 3:
        Input: head = [1,2], n = 1
        Output: [1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	fast := dummy
	slow := dummy

	for i := 0; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

func RunTestRemoveNthFromEnd() {
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
		"case-1": {
			head:   &linkedList1,
			n:      2,
			expect: []int{1, 2, 3, 5},
		},
		"case-2": {
			head:   &linkedList2,
			n:      1,
			expect: []int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		head := removeNthFromEnd(testCase.head, testCase.n)

		result := []int{}
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}

		if !EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
