package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

type P141ListNode struct {
	Val  int
	Next *P141ListNode
}

func hasCycle(head *P141ListNode) bool {
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func mockLinkedList1() *P141ListNode {
	node1 := &P141ListNode{Val: 3}
	node2 := &P141ListNode{Val: 2}
	node3 := &P141ListNode{Val: 0}
	node4 := &P141ListNode{Val: -1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	return node1
}

func mockLinkedList2() *P141ListNode {
	node1 := &P141ListNode{Val: 1}
	node2 := &P141ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1

	return node1
}

func mockLinkedList3() *P141ListNode {
	return &P141ListNode{Val: 1}
}

func RunTestHasCycle() {
	testCases := map[string]struct {
		head   *P141ListNode
		expect bool
	}{
		"case-1": {
			head:   mockLinkedList1(),
			expect: true,
		},
		"case-2": {
			head:   mockLinkedList2(),
			expect: true,
		},
		"case-3": {
			head:   mockLinkedList3(),
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := hasCycle(testCase.head)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
