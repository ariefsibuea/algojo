"""
LeetCode Problem : Two Sum
Topic            : Array, Hash Table
Level            : Easy
URL              : https://leetcode.com/problems/two-sum/description
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
        return nums_map


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