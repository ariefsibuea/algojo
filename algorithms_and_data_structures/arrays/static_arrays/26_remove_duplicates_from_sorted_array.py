"""
LeetCode Problem : Remove Duplicates from Sorted Array
Level            : Easy
URL              : https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
"""

from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        """
        Removes duplicates from a sorted array in-place and returns the number of unique elements.
        This function modifies the input array `nums` such that the first `k` elements of `nums`
        contain the unique elements in sorted order, where `k` is the number of unique elements.
        The remaining elements in the array beyond the first `k` elements are not guaranteed to
        be in any specific order.
        Args:
            nums (List[int]): A list of integers sorted in non-decreasing order.
        Returns:
            int: The number of unique elements in the array after removing duplicates.
        Example:
            nums = [1, 1, 2]
            result = removeDuplicates(nums)
            # nums becomes [1, 2, ...] (remaining elements are unspecified)
            # result is 2
        """

        if len(nums) == 0 or len(nums) == 1:
            return len(nums)

        numUnique = 1
        for counter in range(1, len(nums)):
            if nums[counter] != nums[numUnique - 1]:
                nums[numUnique] = nums[counter]
                numUnique = numUnique + 1

        return numUnique


def run_tests():
    inputs = {
        "case_1": [[1, 1, 2]],
        "case_2": [[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]],
        "case_3": [[1]],
    }
    outputs = {"case_1": 2, "case_2": 5, "case_3": 1}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.removeDuplicates(input[0])
        assert (
            result == outputs[case]
        ), f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
