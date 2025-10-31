from typing import List


def two_sum_problem(nums: List[int], target: int):
    """Finds indices of the two numbers in a list that add up to a specific target.

    Problem:
        Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they
        add up to `target`. Assume that each input would have exactly one solution, and you may not use the same element
        twice.

    Solution:
        Uses a hash map to store the value and its index as we iterate through the list. For each element, checks if the
        complement (target - current element) exists in the hash map. If found, returns the indices.

    Args:
        nums (List[int]): List of integers to search.
        target (int): The target sum to find.

    Returns:
        List[int]: A list containing the indices of the two numbers that add up to the target. Returns an empty list if
        no solution is found.

    Time Complexity:
        O(n): Where n is the number of elements in `nums`, since each element is visited at most twice.

    Space Complexity:
        O(n): Due to the extra space used by the hash map to store up to n elements.
    """
    hasmap = {}
    for i in range(len(nums)):
        complement = target - nums[i]
        if complement in hasmap:
            return [hasmap[complement], i]
        hasmap[nums[i]] = i

    return []


if __name__ == "__main__":
    inputs = {"case_1": [[-2, 1, 2, 4, 7, 11], 13]}

    for _, input in inputs.items():
        print(f"input = {input[0]}\noutput = {two_sum_problem(input[0], input[1])}\n")
