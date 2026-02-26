package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SwappingNodesInALinkedList", RunTestSwappingNodesInALinkedList)
}

/*
 * Problem	: Swapping Nodes in a Linked List
 * Topics	: Linked List, Two Pointers
 * Level	: Medium
 * URL		: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
 *
 * Description:
 * 		You are given the head of a singly linked list and an integer k. Your task is to swap the values of the
 * 		kth node from the beginning and the kth node from the end of the list. The list uses 1-indexing, meaning
 * 		the first node is at position 1. After swapping the values at these two positions, return the head of the
 * 		modified linked list.
 *
 * Constraints:
 * 		- The number of nodes in the list is n
 * 		- 1 <= k <= n <= 10^5
 * 		- 0 <= Node.val <= 100
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [1,2,3,4,5], k = 2
 * 		Output: [1,4,3,2,5]
 * 		Explanation: The 2nd node from the beginning has value 2, and the 2nd node from the end has value 4.
 * 		Swapping these values results in [1,4,3,2,5].
 *
 * 		Example 2:
 * 		Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
 * 		Output: [7,9,6,6,8,7,3,0,9,5]
 * 		Explanation: The 5th node from the beginning has value 7, and the 5th node from the end (which is the
 * 		6th from beginning) has value 8. After swapping, we get [7,9,6,6,8,7,3,0,9,5].
 *
 * 		Example 3:
 * 		Input: head = [1], k = 1
 * 		Output: [1]
 * 		Explanation: With only one node, swapping the 1st node from both ends yields the same list.
 */

func swapNodes(head *ListNode, k int) *ListNode {
	return swapNodesSolutions.withTwoPointers(head, k)
}

type swapNodesSolution struct{}

var swapNodesSolutions = swapNodesSolution{}

func (s *swapNodesSolution) withTwoPointers(head *ListNode, k int) *ListNode {
	firstPtr, secondPtr := (*ListNode)(nil), (*ListNode)(nil)
	slow, fast := head, head

	for i := 0; i < k-1 && fast.Next != nil; i++ {
		fast = fast.Next
	}
	firstPtr = fast

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	secondPtr = slow

	firstPtr.Val, secondPtr.Val = secondPtr.Val, firstPtr.Val
	return head
}

func (s *swapNodesSolution) withExtraSpace(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	curr := dummyHead

	nodes := make([]*ListNode, 0)
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	// NOTE: the values are swapped here, not the nodes
	start, end := 0+k, len(nodes)-k
	nodes[start].Val, nodes[end].Val = nodes[end].Val, nodes[start].Val

	return nodes[0].Next
}

func RunTestSwappingNodesInALinkedList() {
	runner.InitMetrics("SwappingNodesInALinkedList")

	testCases := map[string]struct {
		head   *ListNode
		k      int
		expect []int
	}{
		"example-1-basic-swap": {
			head:   buildSwappingNodesInALinkedListExample1(),
			k:      2,
			expect: []int{1, 4, 3, 2, 5},
		},
		"example-2-complex-list": {
			head:   buildSwappingNodesInALinkedListExample2(),
			k:      5,
			expect: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
		},
		"example-3-single-node": {
			head:   buildSwappingNodesInALinkedListExample3(),
			k:      1,
			expect: []int{1},
		},
		"swap-first-and-last": {
			head:   buildSwappingNodesInALinkedListFiveNodes(),
			k:      1,
			expect: []int{5, 2, 3, 4, 1},
		},
		"two-nodes-swap": {
			head:   buildSwappingNodesInALinkedListTwoNodes(),
			k:      1,
			expect: []int{2, 1},
		},
		"k-equals-middle": {
			head:   buildSwappingNodesInALinkedListOddNodes(),
			k:      3,
			expect: []int{1, 2, 3, 4, 5},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		headSlice := func(h *ListNode) []int {
			var s []int
			for ; h != nil; h = h.Next {
				s = append(s, h.Val)
			}
			return s
		}
		format.PrintInput(map[string]interface{}{"head": headSlice(testCase.head), "k": testCase.k})

		result := runner.ExecCountMetrics(swapNodes, testCase.head, testCase.k).(*ListNode)

		arrResult := make([]int, 0)
		current := result
		for current != nil {
			arrResult = append(arrResult, current.Val)
			current = current.Next
		}

		if !cmp.EqualSlices(arrResult, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v", testCase.expect, arrResult)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}

func buildSwappingNodesInALinkedListExample1() *ListNode {
	return &ListNode{
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
}

func buildSwappingNodesInALinkedListExample2() *ListNode {
	return &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 5,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func buildSwappingNodesInALinkedListExample3() *ListNode {
	return &ListNode{
		Val: 1,
	}
}

func buildSwappingNodesInALinkedListFiveNodes() *ListNode {
	return &ListNode{
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
}

func buildSwappingNodesInALinkedListTwoNodes() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
}

func buildSwappingNodesInALinkedListOddNodes() *ListNode {
	return &ListNode{
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
}
