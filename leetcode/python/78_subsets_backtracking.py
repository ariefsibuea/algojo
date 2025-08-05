"""
LeetCode Problem : Subsets
Topic            : Array, Backtracking, Bit Manipulation
Level            : Medium
URL              : https://leetcode.com/problems/subsets
Description      : Given an integer array nums of unique elements, return all possible subsets (the power set). The
        solution set must not contain duplicate subsets. Return the solution in any order.
Examples         :
        Example 1:
        Input: nums = [1,2,3]
        Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

        Example 2:
        Input: nums = [0]
        Output: [[],[0]]

        Example 3:
        Input: nums = []
        Output: [[]]
"""

from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        subset = []
        self.dfs(0, nums, subset, result)
        return result

    def dfs(self, index: int, nums: List[int], subset: List[int], result: List[List[int]]):
        """Performs depth-first search (DFS) with backtracking to generate all possible subsets (the power set) of a given list.

        Args:
            index (int): The current index in the input list `nums` being considered.
            nums (List[int]): The input list of integers for which subsets are to be generated.
            subset (List[int]): The current subset being constructed.
            result (List[List[int]]): The list that accumulates all generated subsets.

        Returns:
            None: The function modifies the `result` list in place to include all subsets.

        Time Complexity:
            O(n * 2^n):
                2^n possible subsets where n is the length of the input array. Each num has 2 possible subsets.
                We spend O(n) time creating a copy of the subset with list(subset).

        Space Complexity:
            O(n * 2^n):
                O(n) for the recursion stack depth.
                O(n * 2ⁿ) for storing all subsets in the result array
        """

        if index >= len(nums):
            result.append(list(subset))
            return

        subset.append(nums[index])
        self.dfs(index + 1, nums, subset, result)

        subset.pop()
        self.dfs(index + 1, nums, subset, result)


def run_tests():
    inputs = {"case_1": [[1, 2, 3]], "case_2": [[0]]}
    outputs = {"case_1": [[1, 2, 3], [1, 2], [1, 3], [1], [2, 3], [2], [3], []], "case_2": [[0], []]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.subsets(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
