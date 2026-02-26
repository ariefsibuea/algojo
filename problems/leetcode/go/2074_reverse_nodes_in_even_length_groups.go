package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("ReverseNodesInEvenLengthGroups", RunTestReverseNodesInEvenLengthGroups)
}

/*
 * Problem	: Reverse Nodes in Even Length Groups
 * Topics	: Linked List
 * Level	: Medium
 * URL		: https://leetcode.com/problems/reverse-nodes-in-even-length-groups/
 *
 * Description:
 * 		You are given the head of a linked list. The nodes are sequentially assigned to non-empty groups whose
 * 		lengths follow the sequence of natural numbers (1, 2, 3, 4, ...). The first group contains 1 node, the
 * 		second group contains 2 nodes, the third group contains 3 nodes, and so on. The last group may have
 * 		fewer nodes than the expected size (at most 1 + the size of the previous group). Your task is to reverse
 * 		the nodes in each group that has an even number of nodes, while keeping groups with odd lengths unchanged.
 * 		Return the head of the modified linked list after processing all groups.
 *
 * Constraints:
 * 		- 1 <= number of nodes <= 10^5
 * 		- 0 <= Node.val <= 10^5
 *
 * Examples:
 * 		Example 1:
 * 		Input: head = [5,2,6,3,9,1,7,3,8,4]
 * 		Output: [5,6,2,3,9,1,4,8,3,7]
 * 		Explanation: Group 1 has 1 node [5] - odd length, no reversal. Group 2 has 2 nodes [2,6] - even length,
 * 		reversed to [6,2]. Group 3 has 3 nodes [3,9,1] - odd length, no reversal. Last group has 4 nodes
 * 		[7,3,8,4] - even length, reversed to [4,8,3,7]. Final list: [5,6,2,3,9,1,4,8,3,7].
 *
 * 		Example 2:
 * 		Input: head = [1,1,0,6]
 * 		Output: [1,0,1,6]
 * 		Explanation: Group 1 has 1 node [1] - odd length, no reversal. Group 2 has 2 nodes [1,0] - even length,
 * 		reversed to [0,1]. Last group has 1 node [6] - odd length, no reversal. Final list: [1,0,1,6].
 *
 * 		Example 3:
 * 		Input: head = [2,1]
 * 		Output: [2,1]
 * 		Explanation: Group 1 has 1 node [2] - odd length, no reversal. Last group has 1 node [1] - odd length,
 * 		no reversal. Final list remains [2,1].
 *
 * 		Example 4:
 * 		Input: head = [8]
 * 		Output: [8]
 * 		Explanation: Only one group with 1 node - odd length, no reversal needed.
 */

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	return reverseEvenLengthGroupsSolutions.withExtraSpace(head)
}

type reverseEvenLengthGroupsSolution struct{}

var reverseEvenLengthGroupsSolutions = reverseEvenLengthGroupsSolution{}

func (s *reverseEvenLengthGroupsSolution) withExtraSpace(head *ListNode) *ListNode {
	// store nodes into slice
	nodes := make([]*ListNode, 0, 1e5/2)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	// tailor the nodes' order, not including the direction
	groupSize := 1
	currIndex := 0
	for currIndex < len(nodes) {
		// find the actual size
		actualSize := min(groupSize, len(nodes)-currIndex)

		// reverse the nodes if actual size is even
		if actualSize%2 == 0 {
			start, end := currIndex, currIndex+actualSize-1
			for start < end {
				nodes[start], nodes[end] = nodes[end], nodes[start]
				start++
				end--
			}
		}

		currIndex += actualSize
		groupSize++
	}

	// tailor the nodes' direction
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil

	return nodes[0]
}

func RunTestReverseNodesInEvenLengthGroups() {
	runner.InitMetrics("ReverseNodesInEvenLengthGroups")

	testCases := map[string]struct {
		head   *ListNode
		expect []int
	}{
		"example-1-ten-nodes": {
			head:   buildReverseNodesInEvenLengthGroupsExample1(),
			expect: []int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7},
		},
		"example-2-four-nodes": {
			head:   buildReverseNodesInEvenLengthGroupsExample2(),
			expect: []int{1, 0, 1, 6},
		},
		"example-3-two-nodes": {
			head:   buildReverseNodesInEvenLengthGroupsExample3(),
			expect: []int{2, 1},
		},
		"example-4-single-node": {
			head:   buildReverseNodesInEvenLengthGroupsExample4(),
			expect: []int{8},
		},
		"all-odd-groups": {
			head:   buildReverseNodesInEvenLengthGroupsAllOddGroups(),
			expect: []int{1, 3, 2, 5, 4},
		},
		"last-group-even": {
			head:   buildReverseNodesInEvenLengthGroupsLastGroupEven(),
			expect: []int{1, 3, 2, 4},
		},
		"last-group-larger": {
			head:   buildReverseNodesInEvenLengthGroupsLastGroupLarger(),
			expect: []int{1, 3, 2, 5, 4},
		},
		"only-one-group": {
			head:   buildReverseNodesInEvenLengthGroupsOnlyOneGroup(),
			expect: []int{10},
		},
		"two-even-groups": {
			head:   buildReverseNodesInEvenLengthGroupsTwoEvenGroups(),
			expect: []int{1, 3, 2, 4},
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
		format.PrintInput(map[string]interface{}{"head": headSlice(testCase.head)})

		result := runner.ExecCountMetrics(reverseEvenLengthGroups, testCase.head).(*ListNode)

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

func buildReverseNodesInEvenLengthGroupsExample1() *ListNode {
	return NewListFromSlice([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})
}

func buildReverseNodesInEvenLengthGroupsExample2() *ListNode {
	return NewListFromSlice([]int{1, 1, 0, 6})
}

func buildReverseNodesInEvenLengthGroupsExample3() *ListNode {
	return NewListFromSlice([]int{2, 1})
}

func buildReverseNodesInEvenLengthGroupsExample4() *ListNode {
	return NewListFromSlice([]int{8})
}

func buildReverseNodesInEvenLengthGroupsAllOddGroups() *ListNode {
	return NewListFromSlice([]int{1, 2, 3, 4, 5})
}

func buildReverseNodesInEvenLengthGroupsLastGroupEven() *ListNode {
	return NewListFromSlice([]int{1, 2, 3, 4})
}

func buildReverseNodesInEvenLengthGroupsLastGroupLarger() *ListNode {
	return NewListFromSlice([]int{1, 2, 3, 4, 5})
}

func buildReverseNodesInEvenLengthGroupsOnlyOneGroup() *ListNode {
	return NewListFromSlice([]int{10})
}

func buildReverseNodesInEvenLengthGroupsTwoEvenGroups() *ListNode {
	return NewListFromSlice([]int{1, 2, 3, 4})
}
