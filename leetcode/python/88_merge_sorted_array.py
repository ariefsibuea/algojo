"""
LeetCode Problem : Merge Sorted Array
Topic            : Array, Two Pointers, Sorting
Level            : Easy
URL              : https://leetcode.com/problems/merge-sorted-array/
Description      : You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers
        m and n, representing the number of elements in nums1 and nums2 respectively. Merge nums1 and nums2 into a
        single array sorted in non-decreasing order. The final sorted array should not be returned by the function,
        but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the
        first m elements denote the elements that should be merged, and the last n elements are set to 0 and should
        be ignored. nums2 has a length of n.
Examples         :
        Example 1:
        Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
        Output: [1,2,2,3,5,6]
        Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
                     The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

        Example 2:
        Input: nums1 = [1], m = 1, nums2 = [], n = 0
        Output: [1]
        Explanation: The arrays we are merging are [1] and [].
                     The result of the merge is [1].

        Example 3:
        Input: nums1 = [0], m = 0, nums2 = [1], n = 1
        Output: [1]
        Explanation: The arrays we are merging are [] and [1].
                     The result of the merge is [1].
"""

from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """Merges two sorted arrays nums1 and nums2 into a single sorted array in-place.

        Args:
            nums1 (List[int]): The first sorted array with extra space at the end.
            m (int): The number of actual elements in nums1.
            nums2 (List[int]): The second sorted array.
            n (int): The number of elements in nums2.

        Returns:
            None: Modifies nums1 in-place to contain all elements from both arrays in sorted order.

        Time Complexity:
            O(m + n): Each element in nums1 and nums2 is processed once.

        Space Complexity:
            O(1): No extra data structure is used.
        """

        """
        Do not return anything, modify nums1 in-place instead.
        """
        i, j, k = m - 1, n - 1, m + n - 1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1


def run_tests():
    inputs = {"case_1": [[1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3]}
    outputs = {"case_1": [1, 2, 2, 3, 5, 6]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.merge(input[0], input[1], input[2], input[3])
        assert input[0] == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
