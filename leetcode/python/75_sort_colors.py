"""
LeetCode Problem : Sort Colors
Topics           : Array, Two Pointers, Sorting
Level            : Medium
URL              : https://leetcode.com/problems/sort-colors
Description      : Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects
                    of the same color are adjacent, with the colors in the order red, white, and blue. We will use the
                    integers 0, 1, and 2 to represent the color red, white, and blue, respectively. You must solve
                    this problem without using the library's sort function.
Examples         :
                    Example 1:
                    Input: nums = [2,0,2,1,1,0]
                    Output: [0,0,1,1,2,2]

                    Example 2:
                    Input: nums = [2,0,1]
                    Output: [0,1,2]
"""

from typing import List


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        left = 0
        right = len(nums) - 1
        index = 0

        while index <= right:
            if nums[index] == 2:
                nums[right], nums[index] = nums[index], nums[right]
                right -= 1
            elif nums[index] == 0:
                nums[left], nums[index] = nums[index], nums[left]
                left += 1
                index += 1
            else:
                index += 1


def run_tests():
    inputs = {
        "case_1": [
            [2, 0, 2, 1, 1, 0],
        ],
        "case_2": [
            [2, 0, 1],
        ],
        "case_3": [
            [1, 0, 2],
        ],
    }
    outputs = {
        "case_1": [0, 0, 1, 1, 2, 2],
        "case_2": [0, 1, 2],
        "case_3": [0, 1, 2],
    }

    solution = Solution()

    for case, input in inputs.items():
        solution.sortColors(input[0])
        assert input[0] == outputs[case], f"{case}: expected {outputs[case]}, got {input[0]}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
