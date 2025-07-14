"""
LeetCode Problem : 3Sum
Topic            : Array, Two Pointers, Sorting
Level            : Medium
URL              : https://leetcode.com/problems/3sum
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
