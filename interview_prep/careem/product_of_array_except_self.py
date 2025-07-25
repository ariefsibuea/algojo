from typing import List


class ProductOfArrayExceptSelf:
    def __init__(self):
        pass

    def product_except_self(self, nums: List[int]) -> List[int]:
        """Returns a new list where each element is the product of all other elements in the input list.

        Calculates the product of all elements in the input list except for the element at each index, without using
        division. We use two-pass strategy by calculating the left products (product of all elements to the left) and
        the right products (prodict of all elements to the right).

        Args:
            nums (List[int]): A list of integers with length greater than 1.

        Returns:
            List[int]: A list where each element is the product of all elements in nums except for the element at the
            same index.

        Raises:
            ValueError: If nums is None or its length is less than or equal to 1.

        Time Complexity:
            O(n): Where n is the length of nums -> O(n) for the left product + O(n) for the right product.

        Space Complexity:
            O(n): n represents the extra space for output list and auxiliary space for computation.
        """

        n = len(nums)
        if nums is None or n <= 1:
            raise ValueError("length of nums should be greater than 1")

        # Set 1 as initial output result to handle identity multiplication
        output = [1] * n

        # Handle the left product, it will contains product of all elements to left of i
        for i in range(1, n):
            output[i] = output[i - 1] * nums[i - 1]

        # Multiply with the right product
        right_product = 1
        for i in range(n - 1, -1, -1):
            output[i] *= right_product
            right_product *= nums[i]

        return output


if __name__ == "__main__":
    inputs = {"case_1": [[1, 2, 3, 4]], "case_2": [[2, 3, 0, 4]], "case_3": [[0, 0, 1]]}

    outputs = {
        "case_1": [24, 12, 8, 6],
        "case_2": [0, 0, 24, 0],
        "case_3": [0, 0, 0],
    }

    solution = ProductOfArrayExceptSelf()

    for case, input in inputs.items():
        result = solution.product_except_self(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
