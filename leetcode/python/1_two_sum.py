"""
LeetCode Problem : Two Sum
Topic            : Array, Hash Table
Level            : Easy
URL              : https://leetcode.com/problems/two-sum
Description      : Given an array of integers nums and an integer target, return indices of the two numbers such that
    they add up to target. You may assume that each input would have exactly one solution, and you may not use the same
    element twice. You can return the answer in any order.
Examples         :
    Example 1:
    Input: nums = [2,7,11,15], target = 9
    Output: [0,1]
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

    Example 2:
    Input: nums = [3,2,4], target = 6
    Output: [1,2]

    Example 3:
    Input: nums = [3,3], target = 6
    Output: [0,1]
"""

from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """Finds two numbers in the given list that add up to the target and returns their indices.

        Args:
            nums (List[int]): List of integers to search through.
            target (int): Target sum to find.

        Returns:
            List[int]: Indices of the two numbers that add up to target.

        Time Complexity:
            O(n): Each element is visited once in a single pass through the list.

        Space Complexity:
            O(n): In the worst case, all elements are stored in the hash table.
        """
        nums_map = {}
        for i, num in enumerate(nums):
            complement = target - num
            if complement in nums_map:
                return [nums_map[complement], i]
            nums_map[num] = i
        return []


def run_tests():
    inputs = {
        "case_1": [[2, 7, 11, 15], 9],
        "case_2": [[3, 2, 4], 6],
        "case_3": [[3, 3], 6],
    }
    outputs = {"case_1": [0, 1], "case_2": [1, 2], "case_3": [0, 1]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.twoSum(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
