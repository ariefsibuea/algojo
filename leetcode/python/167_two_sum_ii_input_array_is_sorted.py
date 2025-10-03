"""
LeetCode Problem : Two Sum II - Input Array is Sorted
Topics           : Array, Two Pointers, Binary Search
Level            : Medium
URL              : https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
Description      : Given a 1-indexed array sorted in non-decreasing order, you must find the two numbers that add up
                    to a given target. Each input guarantees exactly one solution, and you cannot reuse the same
                    element twice. Return the pair of indices as 1-based positions.
Examples         :
                    Example 1:
                    Input: numbers = [2,7,11,15], target = 9
                    Output: [1,2]
                    Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

                    Example 2:
                    Input: numbers = [2,3,4], target = 6
                    Output: [1,3]
                    Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

                    Example 3:
                    Input: numbers = [-1,0], target = -1
                    Output: [1,2]
                    Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
"""

from typing import List


class Solution:
    def twoSumII(self, numbers: List[int], target: int) -> List[int]:
        left, right = 0, len(numbers) - 1
        sumResult = 0

        while left < right:
            sumResult = numbers[left] + numbers[right]
            if sumResult == target:
                return [left + 1, right + 1]
            elif sumResult > target:
                right -= 1
            else:
                left += 1

        return []


def run_tests():
    inputs = {
        "case_1": [
            [2, 7, 11, 15],
            9,
        ],
        "case_2": [
            [2, 3, 4],
            6,
        ],
        "case_3": [
            [-1, 0],
            -1,
        ],
    }
    outputs = {
        "case_1": [1, 2],
        "case_2": [1, 3],
        "case_3": [1, 2],
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.twoSumII(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
