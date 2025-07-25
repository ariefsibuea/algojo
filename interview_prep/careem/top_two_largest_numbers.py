from typing import List


class TopTwoLargestNumbers:
    def __init__(self):
        pass

    def solve_by_single_pass_tracking(self, nums: List[int]) -> List[int]:
        """Finds the two largest numbers in a list using a single pass with constant space tracking.

        Track the largest and the second largest values seen during traverse elements of the list.

        Args:
            nums (List[int]): A list of integers with at least two elements.

        Returns:
            List[int]: A list containing the largest and second largest numbers in `nums`, in descending order.

        Raises:
            ValueError: If `nums` is None or contains fewer than two elements.

        Time Complexity:
            O(n): Where n is the length of `nums`, as the list is traversed once.

        Space Complexity:
            O(1): As only a constant amount of extra space is used.
        """

        n = len(nums)
        if nums is None or n < 2:
            raise ValueError("'nums' must have at least 2 elements")

        largest = second_largest = float("-inf")

        for num in nums:
            if num > largest:
                second_largest = largest
                largest = num
            elif num > second_largest:
                second_largest = num

        return [largest, second_largest]


if __name__ == "__main__":
    inputs = {
        "case_1": [[3, 1, 4, 1, 5, 9]],
        "case_2": [[10, 10, 10]],
        "case_3": [[0, 5]],
    }

    outputs = {
        "case_1": [9, 5],
        "case_2": [10, 10],
        "case_3": [5, 0],
    }

    solution = TopTwoLargestNumbers()

    for case, input in inputs.items():
        result = solution.solve_by_single_pass_tracking(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
