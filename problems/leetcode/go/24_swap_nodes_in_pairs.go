package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("SwapNodesInPairs", RunTestSwapNodesInPairs)
}

/*
 * Problem	: Swap Nodes in Pairs
 * Topics	: Linked List, Recursion
 * Level	: Medium
 * URL		: https://leetcode.com/problems/swap-nodes-in-pairs/
 *
 * Description:
 * 		Given a linked list, swap every two adjacent nodes and return its head. You must solve the
 * 		problem without modifying the values in the list's nodes (i.e., only nodes themselves may be
 * 		changed). This means you cannot change the node values but must rearrange the node pointers
 * 		to achieve the swapping effect.
 *
 * Constraints:
 * 		- The number of nodes in the list is in the range [0, 100]
 * 		- 0 <= Node.val <= 100
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [1,2,3,4]
 * 		Output: [2,1,4,3]
 * 		Explanation: Nodes 1 and 2 are swapped to become [2,1], nodes 3 and 4 are swapped to become
 * 		[4,3], resulting in [2,1,4,3].
 *
 * 		Example 2:
 * 		Input: head = []
 * 		Output: []
 * 		Explanation: An empty list remains empty after swapping.
 *
 * 		Example 3:
 * 		Input: head = [1]
 * 		Output: [1]
 * 		Explanation: With only one node, there is no pair to swap, so the list remains unchanged.
 *
 * 		Example 4:
 * 		Input: head = [1,2,3]
 * 		Output: [2,1,3]
 * 		Explanation: Only the first pair (1 and 2) can be swapped, leaving node 3 at the end since
 * 		there is no partner to pair it with.
 */

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	curr := prev.Next

	for curr != nil && curr.Next != nil {
		next := curr.Next
		curr.Next = next.Next
		next.Next = curr
		prev.Next = next
		prev = curr
		curr = curr.Next
	}

	return dummyHead.Next
}

func RunTestSwapNodesInPairs() {
	runner.InitMetrics("SwapNodesInPairs")

	testCases := map[string]struct {
		head   *ListNode
		expect []int
	}{
		"example-1-four-nodes": {
			head:   buildSwapNodesInPairsFourNodes(),
			expect: []int{2, 1, 4, 3},
		},
		"example-2-empty-list": {
			head:   nil,
			expect: []int{},
		},
		"example-3-single-node": {
			head:   buildSwapNodesInPairsSingleNode(),
			expect: []int{1},
		},
		"example-4-three-nodes": {
			head:   buildSwapNodesInPairsThreeNodes(),
			expect: []int{2, 1, 3},
		},
		"two-nodes": {
			head:   buildSwapNodesInPairsTwoNodes(),
			expect: []int{2, 1},
		},
		"six-nodes": {
			head:   buildSwapNodesInPairsSixNodes(),
			expect: []int{2, 1, 4, 3, 6, 5},
		},
		"odd-number-of-nodes-five": {
			head:   buildSwapNodesInPairsFiveNodes(),
			expect: []int{2, 1, 4, 3, 5},
		},
	}

	var passedCount uint16 = 0

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]interface{}{"head": testCase.head})

		result := runner.ExecCountMetrics(swapPairs, testCase.head).(*ListNode)
		resultSlice := listNodeToSlice(result)

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

func buildSwapNodesInPairsFourNodes() *ListNode {
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

func buildSwapNodesInPairsSingleNode() *ListNode {
	return &ListNode{
		Val: 1,
	}
}

func buildSwapNodesInPairsThreeNodes() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
}

func buildSwapNodesInPairsTwoNodes() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
}

func buildSwapNodesInPairsSixNodes() *ListNode {
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
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
}

func buildSwapNodesInPairsFiveNodes() *ListNode {
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

func listNodeToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
