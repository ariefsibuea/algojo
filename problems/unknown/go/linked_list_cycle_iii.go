package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Linked List Cycle III
 * Topics           : Linked List, Two Pointers
 * Level            : Medium
 * URL              : -
 * Description      : Given the head of a linked list, determine the length of the cycle present in the linked list. If
 * 					  there is no cycle, return 0.
 * 					  A cycle exists in a linked list if there is some node in the list that can be reached again by
 * 					  continuously following the Next pointer.
 * Constraints      :
 * 					  - The number of nodes in the list is in the range [0, 10^4]
 * 					  - −10^5 <= Node.Value <= 10^5
 * Examples         : -
 */

func countCycleLength(head *ListNode) int {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			cycleLength := 1
			fast = fast.Next

			for fast != slow {
				cycleLength++
				fast = fast.Next
			}

			return cycleLength
		}
	}

	return 0
}

func RunTestLinkedListCycleIII() {
	testCases := map[string]struct {
		head   *ListNode
		expect int
	}{
		"no-cycle": {
			head:   inputTestLinkedListCycleIII_noCycle(),
			expect: 0,
		},
		"cycle-at-end": {
			head:   inputTestLinkedListCycleIII_cycleAtEnd(),
			expect: 3,
		},
		"full-cycle": {
			head:   inputTestLinkedListCycleIII_fullCycle(),
			expect: 2,
		},
		"single-node-cycle": {
			head:   inputTestLinkedListCycleIII_singleNodeCycle(),
			expect: 1,
		},
		"single-node-no-cycle": {
			head:   inputTestLinkedListCycleIII_singleNodeNoCycle(),
			expect: 0,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := countCycleLength(testCase.head)
		format.PrintInput(map[string]interface{}{"head": *testCase.head})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			continue
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func inputTestLinkedListCycleIII_noCycle() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
}

func inputTestLinkedListCycleIII_cycleAtEnd() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Cycle: 2 -> 3 -> 4 -> 2 (length 3)
	return node1
}

func inputTestLinkedListCycleIII_fullCycle() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1 // Cycle: 1 -> 2 -> 1 (length 2)
	return node1
}

func inputTestLinkedListCycleIII_singleNodeCycle() *ListNode {
	node1 := &ListNode{Val: 1}
	node1.Next = node1 // Cycle: 1 -> 1 (length 1)
	return node1
}

func inputTestLinkedListCycleIII_singleNodeNoCycle() *ListNode {
	return &ListNode{Val: 1}
}
