/* Middle of the Linked List
Source		: https://leetcode.com/problems/middle-of-the-linked-list/
Level		: Easy
Description	: Given the head of a singly linked list, return the middle node of the linked list. If there are two middle
			nodes, return the second middle node.

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

package leetcode

// MiddleNode implements two pointers technique to find middle of linked list
func (soln Solution) MiddleNode(head *ListNode) *ListNode {
	mid, cur := head, head
	for cur != nil && cur.Next != nil {
		cur = cur.Next.Next
		mid = mid.Next
	}

	return mid
}
