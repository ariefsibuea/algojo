"""
LeetCode Problem : Climbing Stairs
Topic            : Math, Dynamic Programming, Memoization
Level            : Easy
URL              : https://leetcode.com/problems/climbing-stairs/description/
"""

from typing import Any


class Solution:
    def climbStairs(self, n: int) -> int:
        """Calculates the number of distinct ways to climb n stairs, taking either 1 or 2 steps at a time.

        Args:
            n (int): Number of stairs to climb, where 1 <= n <= 45.

        Returns:
            int: Number of distinct ways to climb the stairs.

        Time Complexity:
            O(n): We iterate through the sequence n times.

        Space Complexity:
            O(1): Only constant extra space is used regardless of input size.
        """

        if n == 1:
            return 1

        last_1, last_2 = 1, 1
        for i in range(1, n):
            temp = last_1
            last_1 = last_1 + last_2
            last_2 = temp

        return last_1


def run_tests():
    inputs = {"case_1": [2], "case_2": [3]}
    outputs = {"case_1": 2, "case_2": 3}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.climbStairs(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
