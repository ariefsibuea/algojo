"""
LeetCode Problem : Container With Most Water
Topic            : Array, Two Pointers, Greedy
Level            : Medium
URL              : https://leetcode.com/problems/container-with-most-water
Description      : You are given an integer array height of length n. There are n vertical lines drawn such that the
        two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form
        a container, such that the container contains the most water. Return the maximum amount of water a container
        can store. Notice that you may not slant the container.
Examples         :
        Example 1:
        Input: height = [1,8,6,2,5,4,8,3,7]
        Output: 49
        Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area
                of water (blue section) the container can contain is 49.

        Example 2:
        Input: height = [1,1]
        Output: 1
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
