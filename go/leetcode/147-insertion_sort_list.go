/* Insertion Sort List
Source		: https://leetcode.com/problems/insertion-sort-list/
Level		: Medium
Description	: Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's
			head.

Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]

Example 2:
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
*/

package leetcode

func (soln Solution) InsertionSortList(head *ListNode) *ListNode {
	root := head
	head = head.Next
	root.Next = nil

	for head != nil {
		if head.Val < root.Val {
			temp := root
			root = head
			head = head.Next
			root.Next = temp
			continue
		}

		currentNode := root
		for {
			if currentNode.Next == nil || currentNode.Next.Val > head.Val {
				temp := currentNode.Next
				currentNode.Next = head
				head = head.Next
				currentNode.Next.Next = temp
				break
			}
			currentNode = currentNode.Next
		}
	}

	return root
}
