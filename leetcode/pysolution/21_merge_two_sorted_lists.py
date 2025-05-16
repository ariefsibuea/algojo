"""
LeetCode Problem : Merge Two Sorted Lists
Topic            : Linked List, Recursion
Level            : Easy
URL              : https://leetcode.com/problems/merge-two-sorted-lists/description/
"""

from typing import Any, Optional


# Default definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        """
        Merge two sorted linked lists and return it as a sorted list.
        Args:
            list1 (Optional[ListNode]): Head of first sorted linked list
            list2 (Optional[ListNode]): Head of second sorted linked list
        Returns:
            Optional[ListNode]: Head of merged sorted linked list
        Example:
            Input: list1 = [1,2,4], list2 = [1,3,4]
            Output: [1,1,2,3,4,4]
        Solution:
            Iterative with Dummy Node
        Time Complexity:
            O(n+m) where n and m are lengths of the input lists
        Space Complexity:
            O(1) as only pointers (dummy and current) are used
        """

        dummy = ListNode(-1)
        current = dummy

        while list1 and list2:
            if list1.val <= list2.val:
                current.next = list1
                list1 = list1.next
            else:
                current.next = list2
                list2 = list2.next
            current = current.next

        current.next = list1 if list1 else list2
        return dummy.next


def run_tests():
    linked_list_case_1 = [
        ListNode(val=1, next=ListNode(val=2, next=ListNode(val=4))),
        ListNode(val=1, next=ListNode(val=3, next=ListNode(val=4))),
    ]

    linked_list_case_2 = [None, None]

    linked_list_case_3 = [None, ListNode(val=0)]

    inputs = {"case_1": linked_list_case_1, "case_2": linked_list_case_2, "case_3": linked_list_case_3}
    outputs = {"case_1": [1, 1, 2, 3, 4, 4], "case_2": [], "case_3": [0]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.mergeTwoLists(input[0], input[1])

        array_result = []
        while result:
            array_result.append(result.val)
            result = result.next
        assert array_result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
