/* Remove Nth Node From End of List
Source		: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Level		: Medium
Description	: Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]
*/

package leetcode

// RemoveNthFromEnd implements two pointers technique to solve remove nth
// node from the end of the linked list.
func (soln Solution) RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	i := 0
	for i = 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	if i < n-1 {
		return head
	}
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
