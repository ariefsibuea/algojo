from typing import List


def is_array_in_sorted_order(nums: List[int], n: int) -> bool:
    """Checks if the given array is sorted in non-decreasing order using recursion.

    Problem:
        Given an array of integers, determine if the array is sorted in non-decreasing order (each element is not less
        than the previous one).

    Solution:
        The function recursively checks if the last element is greater than or equal to the previous element. If so, it
        continues checking the subarray excluding the last element. The base case is when the array has zero or one
        element, which is always sorted.

    Args:
        nums (List[int]): The list of integers to check.
        n (int): The number of elements in the list to consider (typically len(nums)).

    Returns:
        bool: True if the array is sorted in non-decreasing order, False otherwise.

    Time Complexity:
        O(n): Where n is the number of elements in the array.

    Space Complexity:
        O(n): Due to the recursion stack.
    """
    if n == 0 or n == 1:
        return True

    if nums[n-1] < nums[n-2]:
        return False

    return is_array_in_sorted_order(nums, n-1)

if __name__ == "__main__":
    array = [1,2,3,4,5,6,7]
    result = is_array_in_sorted_order(array, len(array))
    print(f"sorted status = {result}")
