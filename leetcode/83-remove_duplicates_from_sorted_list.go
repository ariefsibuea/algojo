/* Remove Duplicates from Sorted List
Source		: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
Level		: Easy
Description	: Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
			Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head
	next := currentNode.Next

	for next != nil {
		if currentNode.Val == next.Val {
			next = next.Next
			currentNode.Next = next
			continue
		}

		temp := next
		next = next.Next
		currentNode = temp
	}

	return head
}
