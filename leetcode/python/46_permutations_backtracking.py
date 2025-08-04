"""
LeetCode Problem : Permutations
Topic            : Array, Backtracking
Level            : Medium
URL              : https://leetcode.com/problems/permutations/description/
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        result = []
        self.backtrack([], nums, [False] * len(nums), result)

        return result

    def backtrack(self, path: List[int], choices: List[int], selected: List[bool], result: List[List[int]]):
        """Uses backtracking to generate all possible permutations of the given choices list by recursively building paths and tracking selected elements.

        Args:
            path (List[int]): The current permutation being constructed.
            choices (List[int]): The list of numbers to permute.
            selected (List[bool]): Boolean list indicating which elements have been used in the current path.
            result (List[List[int]]): Accumulates all complete permutations.

        Returns:
            None: The result list is modified in place to include all permutations.

        Time Complexity:
            O(n * n!):
                O(n!) when generating all permutations.
                O(n) when creating a copy to add to the result.

        Space Complexity:
            O(n * n!):
                O(n!) because we store all permutations in the result list.
                O(n) for the recursion stack and auxiliary arrays.
        """

        if len(path) == len(choices):
            result.append(list(path))
            return

        for i in range(len(choices)):
            if not selected[i]:
                selected[i] = True
                path.append(choices[i])

                self.backtrack(path, choices, selected, result)

                path.pop()
                selected[i] = False


def run_tests():
    inputs = {"case_1": [[1, 2, 3]], "case_2": [[0, 1]], "case_3": [[1]]}
    outputs = {
        "case_1": [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]],
        "case_2": [[0, 1], [1, 0]],
        "case_3": [[1]],
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.permute(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
