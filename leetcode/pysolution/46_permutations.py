"""
LeetCode Problem : Permutations
Topic            : Array, Backtracking
Level            : Medium
URL              : https://leetcode.com/problems/permutations/description/
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        """
        Given an array nums of distinct integers, returns all possible permutations.
        Args:
            nums (List[int]): Array of distinct integers
        Returns:
            List[List[int]]: List of all possible permutations
        Example:
            >>> Solution().permute([1,2,3])
            [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
            >>> Solution().permute([0,1])
            [[0,1], [1,0]]
            >>> Solution().permute([1])
            [[1]]
        Solution:
            Backtracking approach to generate all possible permutations by exploring different paths and tracking used numbers.
        Time Complexity:
            O(n!) where n is the length of input array
        Space Complexity:
            O(n) for recursion stack and tracking used numbers
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
    inputs = {"case_1": [[1, 2, 3]]}
    outputs = {"case_1": [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.permute(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
