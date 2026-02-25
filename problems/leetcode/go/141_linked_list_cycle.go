package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

func init() {
	register("LinkedListCycle", RunTestHasCycle)
}

func hasCycle(head *ListNode) bool {
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

func RunTestHasCycle() {
	testCases := map[string]struct {
		head   *ListNode
		expect bool
	}{
		"case-1": {
			head:   mockInputP141Case1(),
			expect: true,
		},
		"case-2": {
			head:   mockInputP141Case2(),
			expect: true,
		},
		"case-3": {
			head:   mockInputP141Case3(),
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

func mockInputP141Case1() *ListNode {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	return node1
}

func mockInputP141Case2() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1

	return node1
}

func mockInputP141Case3() *ListNode {
	return &ListNode{Val: 1}
}
