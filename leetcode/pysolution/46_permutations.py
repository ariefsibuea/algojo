"""
LeetCode Problem : Permutations
Topic            : Array, Backtracking
Level            : Medium
URL              : https://leetcode.com/problems/permutations/description/
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        """Generates all possible permutations of an array of distinct integers using backtracking.
        
        Args:
            nums (List[int]): Array of distinct integers.
            
        Returns:
            List[List[int]]: List of all possible permutations.
            
        Time Complexity:
            O(n!): Where n is the length of input array, as we generate all permutations.
            
        Space Complexity:
            O(n): Space used for recursion stack and tracking used numbers.
        """

        def backtrack(path: List[int], used: List[bool]):
            if len(path) == len(nums):
                res.append(path.copy())
                return

            for i in range(len(nums)):
                if not used[i]:
                    used[i] = True
                    path.append(nums[i])
                    backtrack(path, used)
                    path.pop()
                    used[i] = False

        res = []
        backtrack([], [False] * len(nums))

        return res


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
