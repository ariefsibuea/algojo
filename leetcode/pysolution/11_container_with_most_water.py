"""
LeetCode Problem : Container With Most Water
Topic            : Array, Two Pointers, Greedy
Level            : Medium
URL              : https://leetcode.com/problems/container-with-most-water/description
"""

from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        """Calculates the maximum amount of water a container can store using the two-pointer approach.

        Args:
            height (List[int]): A list of non-negative integers representing the heights of vertical lines.

        Returns:
            int: The maximum area of water that can be contained.

        Time Complexity:
            O(n): Each element is visited exactly once as the pointers move toward each other.

        Space Complexity:
            O(1): Only using a constant amount of extra variables regardless of input size.
        """
        max_area = 0
        left = 0
        right = len(height) - 1

        while left < right:
            current_width = right - left
            current_height = min(height[right], height[left])
            area = current_width * current_height

            max_area = max(max_area, area)

            if height[left] > height[right]:
                right -= 1
            else:
                left += 1

        return max_area


def run_tests():
    inputs = {"case_1": [[1, 8, 6, 2, 5, 4, 8, 3, 7]], "case_2": [[1, 1]]}
    outputs = {"case_1": 49, "case_2": 1}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.maxArea(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
