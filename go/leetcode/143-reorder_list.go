/* Reorder List
Source		: https://leetcode.com/problems/reorder-list/
Level		: Medium
Description	: You are given the head of a singly linked-list. The list can be represented as:
	L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:
	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
*/

package leetcode

func (soln Solution) ReorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	next, curr := mid.Next, mid.Next.Next
	mid.Next, next.Next = nil, nil
	for curr != nil {
		temp := curr.Next
		curr.Next = next
		next = curr
		curr = temp
	}

	firstHead, secondHead := head, next
	for secondHead != nil {
		temp := firstHead.Next
		firstHead.Next = secondHead
		firstHead = temp

		temp = secondHead.Next
		secondHead.Next = firstHead
		secondHead = temp
	}
}
