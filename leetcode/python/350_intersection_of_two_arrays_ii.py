"""
LeetCode Problem : Intersection of Two Arrays II
Topic            : Array, Hash Table, Two Pointers, Binary Search, Sorting
Level            : Easy
URL              : https://leetcode.com/problems/intersection-of-two-arrays-ii/
Description      : Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in
        the result must appear as many times as it shows in both arrays and you may return the result in any order.
Examples         :
        Example 1:
        Input: nums1 = [1,2,2,1], nums2 = [2,2]
        Output: [2,2]

        Example 2:
        Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
        Output: [4,9]
        Explanation: [9,4] is also accepted.
"""

from typing import List


class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        """Finds the intersection of two arrays, including duplicate elements.

        Args:
            nums1 (List[int]): The first list of integers.
            nums2 (List[int]): The second list of integers.

        Returns:
            List[int]: A list containing the intersection of nums1 and nums2, with duplicates.

        Time Complexity:
            O(n + m): Where n and m are the lengths of nums1 and nums2 respectively.

        Space Complexity:
            O(min(n, m)): Space used for the hash map and result list.
        """

        nums1_map = {}
        for num in nums1:
            nums1_map[num] = nums1_map.get(num, 0) + 1

        result = []

        for num in nums2:
            if nums1_map.get(num, 0) > 0:
                result.append(num)
                nums1_map[num] -= 1

        return result


def run_tests():
    inputs = {"case_1": [[1, 2, 2, 1], [2, 2]], "case_2": [[4, 9, 5], [9, 4, 9, 8, 4]]}
    outputs = {"case_1": [2, 2], "case_2": [9, 4]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.intersect(input[0], input[1])
        assert result == outputs[case], f"Expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
