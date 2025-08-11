"""
LeetCode Problem : Remove Nth Node From End of List
Topic            : Linked List, Two Pointers
Level            : Medium
URL              : https://leetcode.com/problems/remove-nth-node-from-end-of-list
Description      : Given the head of a linked list, remove the nth node from the end of the list and return its head.
Examples         :
        Example 1:
        Input: head = [1,2,3,4,5], n = 2
        Output: [1,2,3,5]

        Example 2:
        Input: head = [1], n = 1
        Output: []

        Example 3:
        Input: head = [1,2], n = 1
        Output: [1]
"""

from typing import Optional


# Default definition of singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        """Remove the nth node from the end of a linked list and return its head.

        Args:
            head (Optional[ListNode]): The head of the linked list
            n (int): The position from the end of the list of the node to be removed (1-based)

        Returns:
            Optional[ListNode]: Head of the modified linked list after removing the nth node from the end

        Solution:
            Two pointers (slow and fast) to find and remove the nth node from the end.

        Time Complexity:
            O(n): n is the length of the linked list

        Space Complexity:
            O(1): only constant extra space is used
        """

        dummy = ListNode(0, head)
        slow = fast = dummy

        # constraint 1 <= n <= linked list size
        for _ in range(n):
            fast = fast.next

        while fast.next:
            fast = fast.next
            slow = slow.next

        slow.next = slow.next.next
        return dummy.next


def run_tests():
    linked_list_case_1 = ListNode(
        val=1, next=ListNode(val=2, next=ListNode(val=3, next=ListNode(val=4, next=ListNode(val=5))))
    )

    linked_list_case_2 = ListNode(val=1)

    inputs = {"case_1": [linked_list_case_1, 2], "case_2": [linked_list_case_2, 1]}
    outputs = {"case_1": [1, 2, 3, 5], "case_2": []}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.removeNthFromEnd(input[0], input[1])

        array_result = []
        while result:
            array_result.append(result.val)
            result = result.next
        assert array_result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
