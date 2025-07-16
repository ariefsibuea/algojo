"""
LeetCode Problem : Permutations
Topic            : Array, Backtracking
Level            : Medium
URL              : https://leetcode.com/problems/permutations
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
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
