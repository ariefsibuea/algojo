package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ReverseLinkedListII", RunTestReverseLinkedListII)
}

/*
 * Problem	: Reverse Linked List II
 * Topics	: Linked List
 * Level	: Medium
 * URL		: https://leetcode.com/problems/reverse-linked-list-ii
 *
 * Description:
 * 		Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes
 * 		of the list from position left to position right, and return the reversed list. Positions are 1-indexed. Nodes
 * 		outside the specified range should remain in their original order.
 *
 * Constraints:
 * 		- The number of nodes in the list is n, where 1 <= n <= 500
 * 		- -500 <= Node.val <= 500
 * 		- 1 <= left <= right <= n
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [1,2,3,4,5], left = 2, right = 4
 * 		Output: [1,4,3,2,5]
 * 		Explanation: Starting with list [1,2,3,4,5], reversing nodes from position 2 to 4
 * 		(2nd through 4th nodes) results in [1,4,3,2,5]. The first and last nodes stay in place.
 *
 * 		Example 2:
 * 		Input: head = [5], left = 1, right = 1
 * 		Output: [5]
 * 		Explanation: When left equals right, no reversal occurs and the list remains unchanged.
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return reverseBetweenSolutions.withTwoPass(head, left, right)
}

type reverseBetweenSolution struct{}

var reverseBetweenSolutions = reverseBetweenSolution{}

func (s *reverseBetweenSolution) withLocalReverse(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next

	}

	return dummyHead.Next
}

func (s *reverseBetweenSolution) withTwoPass(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	nodes := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	start, end := left-1, right-1
	for start < end {
		nodes[start], nodes[end] = nodes[end], nodes[start]
		start++
		end--
	}

	dummyHead := new(ListNode)
	curr = dummyHead
	for i := 0; i < len(nodes); i++ {
		curr.Next = nodes[i]
		curr = curr.Next
	}
	curr.Next = nil

	return dummyHead.Next
}

func RunTestReverseLinkedListII() {
	runner.InitMetrics("ReverseLinkedListII")

	testCases := map[string]struct {
		head   *ListNode
		left   int
		right  int
		expect []int
	}{
		"example-1-basic-reverse": {
			head:   buildReverseLinkedListIIExample1(),
			left:   2,
			right:  4,
			expect: []int{1, 4, 3, 2, 5},
		},
		"example-2-single-node": {
			head:   buildReverseLinkedListIISingleNode(),
			left:   1,
			right:  1,
			expect: []int{5},
		},
		"reverse-entire-list": {
			head:   buildReverseLinkedListIIEntireList(),
			left:   1,
			right:  3,
			expect: []int{3, 2, 1},
		},
		"reverse-at-beginning": {
			head:   buildReverseLinkedListIIAtBeginning(),
			left:   1,
			right:  2,
			expect: []int{2, 1, 3, 4, 5},
		},
		"reverse-at-end": {
			head:   buildReverseLinkedListIIAtEnd(),
			left:   4,
			right:  5,
			expect: []int{1, 2, 3, 5, 4},
		},
		"reverse-middle-two-nodes": {
			head:   buildReverseLinkedListIIMiddleTwo(),
			left:   2,
			right:  3,
			expect: []int{1, 3, 2, 4, 5},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{
			"head":  listToSlice(testCase.head),
			"left":  testCase.left,
			"right": testCase.right,
		})

		result := runner.ExecCountMetrics(reverseBetween, testCase.head, testCase.left, testCase.right).(*ListNode)
		resultSlice := listToSlice(result)

		if !cmp.EqualSlices(resultSlice, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, resultSlice)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	result := []int{head.Val}
	curr := head.Next
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func buildReverseLinkedListIIExample1() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}

func buildReverseLinkedListIISingleNode() *ListNode {
	return &ListNode{Val: 5}
}

func buildReverseLinkedListIIEntireList() *ListNode {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}

func buildReverseLinkedListIIAtBeginning() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}

func buildReverseLinkedListIIAtEnd() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}

func buildReverseLinkedListIIMiddleTwo() *ListNode {
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	return head
}
