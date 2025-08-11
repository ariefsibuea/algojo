"""
LeetCode Problem : Search in Rotated Sorted Array
Topic            : Array, Binary Search
Level            : Medium
URL              : https://leetcode.com/problems/search-in-rotated-sorted-array/
Description      : There is an integer array nums sorted in ascending order (with distinct values). Prior to being
        passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that
        the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
        For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2]. Given the array
        nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if
        it is not in nums.
Examples         :
        Example 1:
        Input: nums = [4,5,6,7,0,1,2], target = 0
        Output: 4

        Example 2:
        Input: nums = [4,5,6,7,0,1,2], target = 3
        Output: -1

        Example 3:
        Input: nums = [1], target = 0
        Output: -1
"""

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        """Searches for a target value in a rotated sorted array and returns its index.

        Args:
            nums (List[int]): The rotated sorted array of integers.
            target (int): The target value to search for.

        Returns:
            int: The index of target if it exists in nums, otherwise -1.

        Time Complexity:
            O(log n): Where n is the length of the array, as each iteration halves the search space.

        Space Complexity:
            O(1): Only constant extra space is used for pointers.
        """

        low, high = 0, len(nums) - 1

        while low <= high:
            mid = low + (high - low) // 2

            if nums[mid] == target:
                return mid
            elif nums[low] <= nums[mid]:
                if nums[low] <= target < nums[mid]:
                    high = mid - 1
                else:
                    low = mid + 1
            else:
                if nums[mid] < target <= nums[high]:
                    low = mid + 1
                else:
                    high = mid - 1

        return -1


def run_tests():
    inputs = {
        "case_1": [[4, 5, 6, 7, 0, 1, 2], 0],
        "case_2": [[4, 5, 6, 7, 0, 1, 2], 3],
        "case_3": [[1], 0],
        "case_4": [[4, 5, 6, 7, 8, 0, 1, 2], 0],
        "case_5": [[1], 2],
    }
    outputs = {"case_1": 4, "case_2": -1, "case_3": -1, "case_4": 5, "case_5": -1}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.search(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
