/* Merge Two Sorted Lists
Source		: https://leetcode.com/problems/merge-two-sorted-lists/
Level		: Easy
Description	: Given the heads of two sorted linked lists list1 and list2. Merge the two lists in a one sorted list. The
			list should be made by splicing together the nodes of the first two lists. Return the head of the merged
			linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
*/

package leetcode

func (soln Solution) MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil && l2 == nil:
		return nil
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var head, currentNode *ListNode
	switch {
	case l1.Val < l2.Val:
		head = l1
		currentNode = head
		l1 = currentNode.Next
	default:
		head = l2
		currentNode = head
		l2 = currentNode.Next
	}

	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			if l1 != nil {
				currentNode.Next = l1
				currentNode = l1
				l1 = currentNode.Next
				continue
			}

			currentNode.Next = l2
			currentNode = l2
			l2 = currentNode.Next
			continue
		}

		if l1.Val < l2.Val {
			currentNode.Next = l1
			currentNode = l1
			l1 = currentNode.Next
			continue
		}

		currentNode.Next = l2
		currentNode = l2
		l2 = currentNode.Next
	}

	return head
}
