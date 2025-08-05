"""
LeetCode Problem : Remove Duplicates from Sorted Array
Topic            : Array, Two Pointers
Level            : Easy
URL              : https://leetcode.com/problems/remove-duplicates-from-sorted-array
Description      : Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that
    each unique element appears only once. The relative order of the elements should be kept the same. Then return the
    number of unique elements in nums.
Examples         :
    Example 1:
    Input: nums = [1,1,2]
    Output: 2, nums = [1,2,_]
    Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.

    Example 2:
    Input: nums = [0,0,1,1,1,2,2,3,3,4]
    Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
    Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4
        respectively.
"""

from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        """Removes duplicates from a sorted list of integers in-place and returns the number of unique elements.

        This solution uses a two-pointer approach to overwrite duplicates and maintain the order of unique elements.

        Args:
            nums (List[int]): A list of sorted integers.

        Returns:
            int: The number of unique elements remaining after duplicates are removed.

        Time Complexity:
            O(n): Where n is the length of the input list.

        Space Complexity:
            O(1): The operation is performed in-place with constant extra space.
        """
        if len(nums) == 0 or len(nums) == 1:
            return len(nums)

        num_unique = 1
        for counter in range(1, len(nums)):
            if nums[counter] != nums[num_unique - 1]:
                nums[num_unique] = nums[counter]
                num_unique = num_unique + 1

        return num_unique


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
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
