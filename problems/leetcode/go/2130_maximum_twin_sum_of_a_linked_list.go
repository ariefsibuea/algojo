package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("MaximumTwinSumOfALinkedList", RunTestMaximumTwinSumOfALinkedList)
}

/*
 * LeetCode Problem : Maximum Twin Sum of a Linked List
 * Topics           : Linked List, Two Pointers, Stack
 * Level            : Medium
 * URL              : https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list
 * Description      : In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known
 * 					  as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.
 * 					  For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2.
 * 					  These are the only nodes with twins for n = 4.
 * 					  The twin sum is defined as the sum of a node and its twin.
 * 					  Given the head of a linked list with even length, return the maximum twin sum of the linked list.
 * Constraints      :
 * 					  - The number of nodes in the list is an even integer in the range [2, 10^5].
 * 					  - 1 <= Node.val <= 10^5
 * Examples         :
 * 					  Example 1:
 * 					  Input: head = [5,4,2,1]
 * 					  Output: 6
 * 					  Explanation:
 * 					  Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
 * 					  There are no other nodes with twins.
 * 					  Thus, the maximum twin sum of the linked list is 6.
 *
 * 					  Example 2:
 * 					  Input: head = [4,2,2,3]
 * 					  Output: 7
 * 					  Explanation:
 * 					  The nodes with twins are:
 * 					  - Node 0 is the twin of node 3, twin sum is 4 + 3 = 7.
 * 					  - Node 1 is the twin of node 2, twin sum is 2 + 2 = 4.
 * 					  The maximum twin sum is max(7, 4) = 7.
 */

func pairSum(head *ListNode) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left, right := (*ListNode)(nil), slow

	for right != nil {
		temp := right.Next
		right.Next = left
		left = right
		right = temp
	}

	maxSum := math.MinInt
	right, left = left, head

	for right != nil {
		maxSum = max(maxSum, left.Val+right.Val)
		left = left.Next
		right = right.Next
	}

	return maxSum
}

func RunTestMaximumTwinSumOfALinkedList() {
	testCases := map[string]struct {
		head   *ListNode
		expect int
	}{
		"two-nodes": {
			head:   inputTestMaximumTwinSumOfALinkedListCaseTwoNodesHead(),
			expect: 9,
		},
		"four-nodes": {
			head:   inputTestMaximumTwinSumOfALinkedListCaseFourNodesHead(),
			expect: 7,
		},
		"six-nodes": {
			head:   inputTestMaximumTwinSumOfALinkedListCaseSixNodesHead(),
			expect: 101,
		},
		"equal-values": {
			head:   inputTestMaximumTwinSumOfALinkedListCaseEqualValuesHead(),
			expect: 2,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := pairSum(testCase.head)
		format.PrintInput(map[string]interface{}{"head": testCase.head})

		if !cmp.EqualNumbers(result, testCase.expect) {
			format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func inputTestMaximumTwinSumOfALinkedListCaseTwoNodesHead() *ListNode {
	return &ListNode{Val: 5, Next: &ListNode{Val: 4}}
}

func inputTestMaximumTwinSumOfALinkedListCaseFourNodesHead() *ListNode {
	return &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
}

func inputTestMaximumTwinSumOfALinkedListCaseSixNodesHead() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 100,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: &ListNode{Val: 1},
					},
				},
			},
		},
	}
}

func inputTestMaximumTwinSumOfALinkedListCaseEqualValuesHead() *ListNode {
	return &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}
}
