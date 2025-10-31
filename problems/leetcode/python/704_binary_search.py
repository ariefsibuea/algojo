"""
LeetCode Problem : Binary Search
Topic            : Array, Binary Search
Level            : Easy
URL              : https://leetcode.com/problems/binary-search/
Description      : Given an array of integers nums which is sorted in ascending order, and an integer target, write a
        function to search target in nums. If target exists, then return its index. Otherwise, return -1. You must
        write an algorithm with O(log n) runtime complexity.
Examples         :
        Example 1:
        Input: nums = [-1,0,3,5,9,12], target = 9
        Output: 4
        Explanation: 9 exists in nums and its index is 4

        Example 2:
        Input: nums = [-1,0,3,5,9,12], target = 2
        Output: -1
        Explanation: 2 does not exist in nums so return -1
"""

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        """Implements binary search to find a target value in a sorted array.

        Args:
            nums (List[int]): Array of integers sorted in ascending order.
            target (int): Target value to search for.

        Returns:
            int: Index if target exists in nums, otherwise -1.

        Time Complexity:
            O(log n): Where n is the length of the nums array.

        Space Complexity:
            O(1): Only constant extra space is used for pointers.
        """

        low, high = 0, len(nums) - 1
        while low <= high:
            mid = low + (high - low) // 2

            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                high = mid - 1
            else:
                low = mid + 1
        return -1


def run_tests():
    inputs = {"case_1": [[-1, 0, 3, 5, 9, 12], 9], "case_2": [[-1, 0, 3, 5, 9, 12], 2]}
    outputs = {"case_1": 4, "case_2": -1}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.search(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
