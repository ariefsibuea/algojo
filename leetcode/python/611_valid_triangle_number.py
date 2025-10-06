"""
LeetCode Problem : Valid Triangle Number
Topics           : Array, Two Pointers, Binary Search, Greedy, Sorting
Level            : Medium
URL              : https://leetcode.com/problems/valid-triangle-number
Description      : Given an integer array nums, return the number of triplets chosen from the array that can make
                    triangles if we take them as side lengths of a triangle.
Examples         :
                    Example 1:
                    Input: nums = [2,2,3,4]
                    Output: 3
                    Explanation: Valid combinations are:
                    2,3,4 (using the first 2)
                    2,3,4 (using the second 2)
                    2,2,3
"""

from typing import List


class Solution:
    def triangleNumber(self, nums: List[int]) -> int:
        nums.sort() # sort asecending

        nums_len = len(nums)
        count = 0

        for i in range(nums_len - 1, -1, -1):
            n = nums[i]
            left = 0
            right = i - 1

            while left < right:
                # NOTE: A triangle is valid if a + b > c, a + c > b, b + c > a. Select the left most number and assume
                # it as one of the side length. Since it is the most length, then the other two are the lowest and the
                # sum result must be greather the selected one.
                if nums[left] + nums[right] > n:
                    count += right - left
                    right -= 1
                else:
                    left += 1

        return count


def run_tests():
    inputs = {
        "case_1": [
            [2,2,3,4]
        ],
        "case_2": [
            [4,2,3,4]
        ],
    }
    outputs = {
        "case_1": 3,
        "case_2": 4,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.triangleNumber(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
