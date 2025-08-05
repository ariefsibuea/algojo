"""
LeetCode Problem : Maximum Subarray
Topic            : Array, Divide and Conquer, Dynamic Programming
Level            : Medium
URL              : https://leetcode.com/problems/maximum-subarray
Description      : Given an integer array nums, find the subarray with the largest sum, and return its sum.
Examples         :
        Example 1:
        Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
        Output: 6
        Explanation: The subarray [4,-1,2,1] has the largest sum 6.

        Example 2:
        Input: nums = [1]
        Output: 1
        Explanation: The subarray [1] has the largest sum 1.

        Example 3:
        Input: nums = [5,4,-1,7,8]
        Output: 23
        Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
"""

from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        """Finds the contiguous subarray within a one-dimensional array of numbers which has the largest sum using Kadane's algorithm.

        Args:
            nums (List[int]): A list of integers.

        Returns:
            int: The largest sum of any contiguous subarray.

        Time Complexity:
            O(n): Each element is processed exactly once with constant-time operations.

        Space Complexity:
            O(1): Only two variables (sum and max_subarray) regardless of input size
        """

        sum = 0
        max_subarray = float("-inf")

        for _, num in enumerate(nums):
            if sum < 0:
                sum = 0
            sum += num

            max_subarray = max(max_subarray, sum)

        return int(max_subarray)


def run_tests():
    inputs = {"case_1": [[-2, 1, -3, 4, -1, 2, 1, -5, 4]], "case_2": [[1]], "case_3": [[5, 4, -1, 7, 8]]}
    outputs = {"case_1": 6, "case_2": 1, "case_3": 23}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.maxSubArray(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
