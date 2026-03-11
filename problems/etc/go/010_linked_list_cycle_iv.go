package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("LinkedListCycleIV", RunTestLinkedListCycleIV)
}

/*
 * Problem 			: Linked List Cycle IV
 * Topics           : Linked List, Two Pointers
 * Level            : Medium
 * URL              : -
 * Description      : Given the head of a singly linked list, implement a function to detect and remove any cycle
 * 					  present in the list. A cycle occurs when a node's Next pointer links back to a previous node,
 * 					  forming a loop within the list.
 * 					  The function must modify the linked list in place, ensuring it remains acyclic while preserving
 * 					  the original node order. If no cycle is found, return the linked list as is.
 * Constraints      :
 * 					  - The number of nodes in the list is in the range [0, 10^4]
 * 					  - −10^5 <= Node.Value <= 10^5
 * Examples         : -
 */

func removeCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasCycle := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return head
	}

	slow = head

	if slow == fast {
		for fast.Next != slow {
			fast = fast.Next
		}
		fast.Next = nil
		return head
	}

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	for fast.Next != slow {
		fast = fast.Next
	}

	fast.Next = nil
	return head
}

func RunTestLinkedListCycleIV() {
	runner.InitMetrics("LinkedListCycleIV")

	testCases := map[string]struct {
		head   *ListNode
		expect []int
	}{
		"empty-list": {
			head:   inputTestLinkedListCycleIVCaseEmptyListHead(),
			expect: []int{},
		},
		"single-node-no-cycle": {
			head:   inputTestLinkedListCycleIVCaseSingleNodeNoCycleHead(),
			expect: []int{1},
		},
		"single-node-cycle": {
			head:   inputTestLinkedListCycleIVCaseSingleNodeCycleHead(),
			expect: []int{1},
		},
		"multi-nodes-no-cycle": {
			head:   inputTestLinkedListCycleIVCaseMultiNodesNoCycleHead(),
			expect: []int{1, 2, 3, 4},
		},
		"multi-nodes-cycle-at-end": {
			head:   inputTestLinkedListCycleIVCaseMultiNodesCycleAtEndHead(),
			expect: []int{1, 2, 3, 4, 5},
		},
		"multi-nodes-cycle-at-start": {
			head:   inputTestLinkedListCycleIVCaseMultiNodesCycleAtStartHead(),
			expect: []int{1, 2, 3, 4, 5},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"head": testCase.head})

		head := runner.ExecCountMetrics(removeCycle, testCase.head).(*ListNode)
		result := []int{}
		curr := head
		for curr != nil {
			result = append(result, curr.Val)
			curr = curr.Next
		}

		if !cmp.EqualSlices(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, result)
			continue
		}
		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

func inputTestLinkedListCycleIVCaseEmptyListHead() *ListNode {
	return nil
}

func inputTestLinkedListCycleIVCaseSingleNodeNoCycleHead() *ListNode {
	return &ListNode{Val: 1}
}

func inputTestLinkedListCycleIVCaseSingleNodeCycleHead() *ListNode {
	node := &ListNode{Val: 1}
	node.Next = node
	return node
}

func inputTestLinkedListCycleIVCaseMultiNodesNoCycleHead() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
}

func inputTestLinkedListCycleIVCaseMultiNodesCycleAtEndHead() *ListNode {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n2

	return n1
}

func inputTestLinkedListCycleIVCaseMultiNodesCycleAtStartHead() *ListNode {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n1

	return n1
}
