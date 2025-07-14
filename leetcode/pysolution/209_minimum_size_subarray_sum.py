"""
LeetCode Problem : Minimum Size Subarray Sum
Topic            : Array, Binary Search, Sliding Window, Prefix Sum
Level            : Medium
URL              : https://leetcode.com/problems/minimum-size-subarray-sum/description
"""

from typing import List


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        """Finds the minimal length of a contiguous subarray of which the sum is at least the given target value.

        Args:
            target (int): The target sum to be achieved by the subarray.
            nums (List[int]): List of positive integers.

        Returns:
            int: The minimal length of such a subarray, or 0 if no such subarray exists.

        Time Complexity:
            O(n): Each element is added to the window once and removed at most.

        Space Complexity:
            O(1): Only use a constant amount of extra space.
        """

        if not nums:
            return 0

        min_length = len(nums) + 1
        left = 0
        sum = 0

        for right in range(len(nums)):
            sum += nums[right]

            while sum >= target:
                min_length = min(min_length, right - left + 1)
                sum -= nums[left]
                left += 1

        return min_length if min_length < len(nums) + 1 else 0


def run_tests():
    inputs = {"case_1": [7, [2, 3, 1, 2, 4, 3]], "case_2": [4, [1, 4, 4]], "case_3": [11, [1, 1, 1, 1, 1, 1, 1, 1]]}
    outputs = {"case_1": 2, "case_2": 1, "case_3": 0}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.minSubArrayLen(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
