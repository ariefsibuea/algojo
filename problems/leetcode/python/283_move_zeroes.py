"""
LeetCode Problem : Move Zeroes
Topics           : Array, Two Pointers
Level            : Easy
URL              : https://leetcode.com/problems/move-zeroes
Description      : Given an integer array nums, move all 0's to the end of it while maintaining the relative order of
                    the non-zero elements. Note that you must do this in-place without making a copy of the array.
Examples         :
                    Example 1:
                    Input: nums = [0,1,0,3,12]
                    Output: [1,3,12,0,0]

                    Example 2:
                    Input: nums = [0]
                    Output: [0]
"""

from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        left = 0
        for right in range(len(nums)):
            if nums[right] != 0:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1


def run_tests():
    inputs = {
        "case_1": [
            [0, 1, 0, 3, 12],
        ],
        "case_2": [
            [0],
        ],
    }
    outputs = {
        "case_1": [1, 3, 12, 0, 0],
        "case_2": [0],
    }

    solution = Solution()

    for case, input in inputs.items():
        solution.moveZeroes(input[0])
        assert input[0] == outputs[case], f"{case}: expected {outputs[case]}, got {input[0]}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
