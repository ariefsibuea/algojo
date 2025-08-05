"""
LeetCode Problem : 3Sum
Topic            : Array, Two Pointers, Sorting
Level            : Medium
URL              : https://leetcode.com/problems/3sum
Description      : Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j,
        i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0. Notice that the solution set must not contain
        duplicate triplets.
Examples         :
        Example 1:
        Input: nums = [-1,0,1,2,-1,-4]
        Output: [[-1,-1,2],[-1,0,1]]
        Explanation: nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0. nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
                nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0. The distinct triplets are [-1,0,1] and [-1,-1,2].

        Example 2:
        Input: nums = [0,1,1]
        Output: []
        Explanation: The only possible triplet sums up to 3, not 0.

        Example 3:
        Input: nums = [0,0,0]
        Output: [[0,0,0]]
        Explanation: The only possible triplet sums up to 0.
"""

from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        """Finds all unique triplets in the list that sum up to zero. The solution combining sort and two-pointers approach.

        Args:
            nums (List[int]): List of integers to search for triplets.

        Returns:
            List[List[int]]: A list of lists, each containing three integers that sum to zero.

        Time Complexity:
            O(n^2):
                Sorting: O(n log n)
                Nested loops: O(n^2) where the outer loop runs n times in the worst case, and for each iteration, the
                            two-pointer loop runs at most n times.

        Space Complexity:
            O(n): Only a constant amount of additional variables are used, the sorting is typically done in-place in Python.
        """

        nums.sort()

        result = []
        index = 0

        for index in range(len(nums) - 2):
            num = nums[index]
            if index > 0 and num == nums[index - 1]:
                index += 1
                continue

            left = index + 1
            right = len(nums) - 1
            while left < right:
                sum = num + nums[left] + nums[right]

                if sum > 0:
                    right -= 1
                elif sum < 0:
                    left += 1
                else:
                    result.append([num, nums[left], nums[right]])
                    left += 1

                    while nums[left] == nums[left - 1] and left < right:
                        left += 1

            index += 1

        return result


def run_tests():
    inputs = {"case_1": [[-1, 0, 1, 2, -1, -4]]}
    outputs = {"case_1": [[-1, -1, 2], [-1, 0, 1]]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.threeSum(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
