from typing import List


def array_advance_game(nums: List[int]) -> bool:
    """Returns true if the last index can be reached, or false otherwise.

    Problem:
        The "array advance game" is a problem where you are given an array of non-negative integers, where each element
        represents your maximum jump length at that position. Starting at the first index, the goal is to determine if
        you can reach the last index by making jumps according to the values in the array.

    Solution:
        The function uses a greedy approach to keep track of the furthest index that can be reached at each step. It
        iterates through the array, updating the furthest reachable index. If at any point the current index exceeds the
        furthest reachable index, it means the end cannot be reached. The function returns True if the last index is
        reachable, otherwise False.

    Args:
        nums (List[int]): A list of non-negative integers representing the maximum jump length from each position.

    Returns:
        bool: True if the last index is reachable from the first index, False otherwise.

    Time Complexity:
        O(n): Where n is the length of the input array.

    Space Complexity:
        O(1): Only a constant amount of extra space is used.
    """
    furthest_reached = 0
    last_index = len(nums) - 1
    index = 0
    while index <= furthest_reached and furthest_reached < last_index:
        furthest_reached = max(furthest_reached, nums[index] + index)
        index += 1

    return furthest_reached >= last_index


if __name__ == "__main__":
    inputs = {"case_1": [[3, 3, 1, 0, 2, 0, 1]], "case_2": [[3, 2, 0, 0, 2, 0, 1]]}

    for _, input in inputs.items():
        result = array_advance_game(input[0])
        print(f"input = {input[0]}\noutput = {result}\n")
