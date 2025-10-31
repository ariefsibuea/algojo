"""
LeetCode Problem : Permutations
Topic            : Array, Backtracking
Level            : Medium
URL              : https://leetcode.com/problems/permutations
Description      : Given an array nums of distinct integers, return all the possible permutations. You can return the
        answer in any order.
Examples         :
        Example 1:
        Input: nums = [1,2,3]
        Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

        Example 2:
        Input: nums = [0,1]
        Output: [[0,1],[1,0]]

        Example 3:
        Input: nums = [1]
        Output: [[1]]
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        """Generates all possible permutations of a list of integers using an iterative insertion approach.
        At each step, the current number is inserted into every possible position of all existing permutations.

        Args:
            nums (List[int]): The list of integers to permute.

        Returns:
            List[List[int]]: A list containing all possible permutations of the input list.

        Time Complexity:
            O(n * n!): Where n is the length of nums, since there are n! permutations and each insertion takes up to O(n) time.

        Space Complexity:
            O(n * n!): All permutations are stored in memory.
        """

        permutations = [[]]

        for num in nums:
            new_permutations = []
            for p in permutations:
                for i in range(len(p) + 1):
                    new_p = p.copy()
                    new_p.insert(i, num)
                    new_permutations.append(new_p)
            permutations = new_permutations

        return permutations


def run_tests():
    inputs = {"case_1": [[1, 2, 3]], "case_2": [[0, 1]], "case_3": [[1]]}
    outputs = {
        "case_1": [[3, 2, 1], [2, 3, 1], [2, 1, 3], [3, 1, 2], [1, 3, 2], [1, 2, 3]],
        "case_2": [[1, 0], [0, 1]],
        "case_3": [[1]],
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.permute(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
