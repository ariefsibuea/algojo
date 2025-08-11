"""
LeetCode Problem : Median of Two Sorted Arrays
Topic            : Array, Binary Search, Divide and Conquer
Level            : Hard
URL              : https://leetcode.com/problems/median-of-two-sorted-arrays
Description      : Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two
        sorted arrays. The overall run time complexity should be O(log (m+n)).
Examples         :
        Example 1:
        Input: nums1 = [1,3], nums2 = [2]
        Output: 2.00000
        Explanation: merged array = [1,2,3] and median is 2.

        Example 2:
        Input: nums1 = [1,2], nums2 = [3,4]
        Output: 2.50000
        Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
"""

from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float | None:
        """Finds the median of two sorted arrays using a binary search approach to partition the arrays efficiently.

        Args:
            nums1 (List[int]): The first sorted array.
            nums2 (List[int]): The second sorted array.

        Returns:
            float: The median value of the combined sorted arrays.

        Time Complexity:
            O(log(min(m, n))): Perform binary search on the smaller array and each comparison takes constant time.

        Space Complexity:
            O(1): Only use a constant amount of extra space regardless of input sizes.
        """

        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        x, y = len(nums1), len(nums2)
        low, high = 0, x
        half = (x + y) // 2

        while low <= high:
            part_x = (high + low) // 2
            part_y = half - part_x

            low_x = float("-inf") if part_x <= 0 else nums1[part_x - 1]
            high_x = float("inf") if part_x >= x else nums1[part_x]
            low_y = float("-inf") if part_y <= 0 else nums2[part_y - 1]
            high_y = float("inf") if part_y >= y else nums2[part_y]

            if low_x <= high_y and low_y <= high_x:
                if (x + y) % 2:
                    return min(high_x, high_y)
                else:
                    return (max(low_x, low_y) + min(high_x, high_y)) / 2
            elif high_x < low_y:
                low = part_x + 1
            else:
                high = part_x - 1


def run_tests():
    inputs = {"case_1": [[1, 3], [2]], "case_2": [[1, 2], [3, 4]]}
    outputs = {"case_1": 2, "case_2": 2.5}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.findMedianSortedArrays(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"


if __name__ == "__main__":
    run_tests()
