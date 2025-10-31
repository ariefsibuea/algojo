"""
LeetCode Problem : Trapping Rain Water
Topics           : Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack
Level            : Hard
URL              : https://leetcode.com/problems/trapping-rain-water
Description      : Given n non-negative integers representing an elevation map where the width of each bar is 1,
                    compute how much water it can trap after raining.
Examples         :
                    Example 1:
                    Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
                    Output: 6
                    Explanation: The above elevation map (black section) is represented by array
                    [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

                    Example 2:
                    Input: height = [4,2,0,3,2,5]
                    Output: 9
"""

from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        left, right = 0, len(height) - 1
        max_left, max_right = height[left], height[right]

        trapped_count = 0
        while left < right:
            if max_left < max_right:
                left += 1
                if max_left < height[left]:
                    max_left = height[left]
                else:
                    trapped_count += max_left - height[left]
            else:
                right -= 1
                if max_right < height[right]:
                    max_right = height[right]
                else:
                    trapped_count += max_right - height[right]

        return trapped_count


def run_tests():
    inputs = {
        "case_1": [
            [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1],
        ],
        "case_2": [
            [4, 2, 0, 3, 2, 5],
        ],
    }
    outputs = {
        "case_1": 6,
        "case_2": 9,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.trap(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
