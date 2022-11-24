/* Convert Binary Number in a Linked List to Integer
Source		: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
Level		: Easy
Description	:
	Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0
	or 1. The linked list holds the binary representation of a number.
	Return the decimal value of the number in the linked list.
	The most significant bit is at the head of the linked list.

Example 1:
Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10

Example 2:
Input: head = [0]
Output: 0
*/

package leetcode

func GetDecimalValue(head *ListNode) int {
	decimalNum := head.Val
	for head.Next != nil {
		decimalNum = (decimalNum * 2) + head.Next.Val
		head = head.Next
	}
	return decimalNum
}
