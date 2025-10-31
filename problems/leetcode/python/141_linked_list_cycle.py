"""
LeetCode Problem : Linked List Cycle
Topic            : Hash Table, Linked List, Two Pointers
Level            : Easy
URL              : https://leetcode.com/problems/linked-list-cycle/
Description      : Given head, the head of a linked list, determine if the linked list has a cycle in it. There is a
        cycle in a linked list if there is some node in the list that can be reached again by continuously following the
        next pointer. Return true if there is a cycle in the linked list. Otherwise, return false.
Examples         :
        Example 1:
        Input: head = [3,2,0,-4], pos = 1
        Output: true
        Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

        Example 2:
        Input: head = [1,2], pos = 0
        Output: true
        Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

        Example 3:
        Input: head = [1], pos = -1
        Output: false
        Explanation: There is no cycle in the linked list.
"""

from typing import Optional


# Default definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        """Determines if a linked list contains a cycle using Floyd's Cycle Finding Algorithm.

        Args:
            head (Optional[ListNode]): Head of the linked list.

        Returns:
            bool: True if the linked list has a cycle, False otherwise.

        Time Complexity:
            O(n): Where n is the number of nodes in the linked list.

        Space Complexity:
            O(1): Only constant extra space is used for two pointers.
        """

        slow, fast = head, head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

            if slow == fast:
                return True

        return False


def run_tests():
    node_1_1 = ListNode(3)
    node_1_2 = ListNode(2)
    node_1_3 = ListNode(0)
    node_1_4 = ListNode(-4)
    node_1_1.next = node_1_2
    node_1_2.next = node_1_3
    node_1_3.next = node_1_4
    node_1_4.next = node_1_2
    linked_list_case_1 = node_1_1

    node_2_1 = ListNode(1)
    node_2_2 = ListNode(2)
    node_2_1.next = node_2_2
    node_2_2.next = node_2_1
    linked_list_case_2 = node_2_1

    linked_list_case_3 = ListNode(1)

    inputs = {"case_1": [linked_list_case_1], "case_2": [linked_list_case_2], "case_3": [linked_list_case_3]}
    outputs = {"case_1": True, "case_2": True, "case_3": False}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.hasCycle(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
