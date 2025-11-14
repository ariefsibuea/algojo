package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Add Two Numbers
 * Topics           : Linked List, Math, Recursion
 * Level            : Medium
 * URL              : https://leetcode.com/problems/add-two-numbers
 * Description      : Given two non-empty linked lists representing two non-negative integers. The digits are stored
 * 					in reverse order, and each of their nodes contains a single digit. Add the two numbers and return
 * 					the sum as a linked list. Assume the two numbers do not contain any leading zero, except the
 * 					number 0 itself.
 * Examples         :
 * 					Example 1:
 * 					Input: l1 = [2,4,3], l2 = [5,6,4]
 * 					Output: [7,0,8]
 * 					Explanation: 342 + 465 = 807.
 *
 * 					Example 2:
 * 					Input: l1 = [0], l2 = [0]
 * 					Output: [0]
 *
 * 					Example 3:
 * 					Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 					Output: [8,9,9,9,0,0,0,1]
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead = &ListNode{}
	var current = dummyHead
	var carry, sum = 0, 0
	var num1, num2 = 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		num1 = getListNodeValue(l1)
		num2 = getListNodeValue(l2)

		sum = (carry + num1 + num2) % 10
		carry = (carry + num1 + num2) / 10

		newNode := &ListNode{Val: sum}
		current.Next = newNode
		current = newNode

		l1 = nextNode(l1)
		l2 = nextNode(l2)
	}

	return dummyHead.Next
}

func RunTestAddTwoNumbers() {
	testCases := map[string]struct {
		l1     *ListNode
		l2     *ListNode
		expect []int
	}{
		"case-1": {
			l1:     mockInputP2Case1()[0],
			l2:     mockInputP2Case1()[1],
			expect: []int{7, 0, 8},
		},
		"case-2": {
			l1:     mockInputP2Case2()[0],
			l2:     mockInputP2Case2()[1],
			expect: []int{0},
		},
		"case-3": {
			l1:     mockInputP2Case3()[0],
			l2:     mockInputP2Case3()[1],
			expect: []int{8, 9, 0, 1},
		},
		"case-4": {
			l1:     mockInputP2Case3()[0],
			l2:     mockInputP2Case3()[1],
			expect: []int{8, 9, 0, 1},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := addTwoNumbers(testCase.l1, testCase.l2)

		arrResult := make([]int, 0)
		for result != nil {
			arrResult = append(arrResult, result.Val)
			result = result.Next
		}

		if !cmp.EqualSlices(arrResult, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP2Case1() []*ListNode {
	head1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}

	head2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}

	return []*ListNode{
		head1,
		head2,
	}
}

func mockInputP2Case2() []*ListNode {
	head1 := &ListNode{Val: 0}
	head2 := &ListNode{Val: 0}

	return []*ListNode{
		head1,
		head2,
	}
}

func mockInputP2Case3() []*ListNode {
	head1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{Val: 9},
		},
	}

	head2 := &ListNode{
		Val:  9,
		Next: &ListNode{Val: 9},
	}

	return []*ListNode{
		head1,
		head2,
	}
}

func mockInputP2Case4() []*ListNode {
	head1 := &ListNode{
		Val:  9,
		Next: &ListNode{Val: 9},
	}

	head2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{Val: 9},
		},
	}

	return []*ListNode{
		head1,
		head2,
	}
}
