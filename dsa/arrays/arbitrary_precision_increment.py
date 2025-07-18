from typing import List


def arbitrary_precision_increment(nums: List[int]) -> List[int]:
    """Increments an integer represented as a list of its digits by one.

    Problem:
        Given a non-empty list of non-negative integers where each element represents a single digit of a non-negative
        integer (with the most significant digit at the start of the list), increment the integer by one and return the
        resulting list of digits.

    Solution:
        The function starts by incrementing the last digit. It then iterates backwards through the list, handling any
        carry that results from a digit becoming 10 (setting it to 0 and incrementing the previous digit). If the most
        significant digit becomes 10 after all carries are processed, it is set to 1 and a new 0 is appended to the end
        of the list to represent the new least significant digit.

    Args:
        nums (List[int]): A list of integers where each integer is a digit (0-9) of a non-negative integer.

    Returns:
        List[int]: A list of integers representing the incremented value.

    Time Complexity:
        O(n): Where n is the number of digits in the input list.

    Space Complexity:
        O(1): The operation is performed in-place except for the possible addition of one digit in the case of an
        overflow.
    """
    nums[-1] += 1
    for i in range(len(nums)-1, 0, -1):
        if nums[i] != 10:
            break
        nums[i] = 0
        nums[i-1] += 1
    if nums[0] == 10:
        nums[0] = 1
        nums.append(0)
    return nums

if __name__ == "__main__":
    inputs = {"case_1": [[1, 4, 9]], "case_2": [[9, 9, 9]]}

    for _, input in inputs.items():
        print(f"input = {input[0]}\noutput = {arbitrary_precision_increment(input[0])}\n")
