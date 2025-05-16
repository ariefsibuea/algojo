"""
LeetCode Problem : Binary Search
Topic            : Array, Binary Search
Level            : Easy
URL              : https://leetcode.com/problems/binary-search/description/
"""

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        """Given an array of integers nums which is sorted in ascending order, and an integer target,
        implements binary search algorithm to search target in nums.
        Args:
            nums (List[int]): Array of integers sorted in ascending order
            target (int): Target value to search for
        Returns:
            int: Index if target exists in nums, otherwise return -1
        Example:
            >>> solution = Solution()
            >>> solution.search([-1,0,3,5,9,12], 9)
            4
            >>> solution.search([-1,0,3,5,9,12], 2)
            -1
        Solution:
            Iterative binary search
        Time Complexity:
            O(log n) where n is the length of nums array.
        Space Complexity:
            O(1) where uses only constant space for pointers (low, high, mid).
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
