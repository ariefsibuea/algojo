"""
LeetCode Problem : Reverse Linked List
Topic            : Linked List, Recursion
Level            : Easy
URL              : https://leetcode.com/problems/reverse-linked-list/description/
"""

from typing import Optional


# Default definition of singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        """Reverses a singly linked list using an iterative approach.
        
        Args:
            head (Optional[ListNode]): The head node of the singly linked list.
            
        Returns:
            Optional[ListNode]: The new head node of the reversed linked list.
            
        Time Complexity:
            O(n): Where n is the number of nodes in the list, as each node is visited once.
            
        Space Complexity:
            O(1): Only constant extra space is used for pointers.
        """

        previous = None
        current = head

        while current:
            next = current.next
            current.next = previous
            previous = current
            current = next

        return previous


def run_tests():
    linked_list_case_1 = ListNode(
        val=1, next=ListNode(val=2, next=ListNode(val=3, next=ListNode(val=4, next=ListNode(val=5, next=None))))
    )

    linked_list_case_2 = ListNode(val=1, next=ListNode(val=2, next=None))

    inputs = {"case_1": [linked_list_case_1], "case_2": [linked_list_case_2]}
    outputs = {"case_1": [5, 4, 3, 2, 1], "case_2": [2, 1]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.reverseList(input[0])

        array_result = []
        while result.next:
            array_result.append(result.val)
            result = result.next
        array_result.append(result.val)

        assert array_result == outputs[case], f"{case}: expected {outputs[case]}, got {array_result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
