class ValidateParentheses:
    def __init__(self):
        pass

    def is_valid_parentheses(self, parentheses: str) -> bool:
        """Checks if the input string of parentheses is valid (all opening brackets are properly closed and nested).

        Use a stack data structure to track opening brackets. When we encounter a closing bracket, we check if it
        matches the most recent opening bracket by popping from the stack.

        Args:
            parentheses (str): The string containing parentheses to validate.

        Returns:
            bool: True if the parentheses are valid, False otherwise.

        Raises:
            ValueError: If the input string is None.

        Time Complexity:
            O(n): Where n is the length of the input string.

        Space Complexity:
            O(n): Due to the use of a stack to store opening brackets.
        """

        if parentheses is None:
            raise ValueError("parentheses is empty")

        bracket_pair = {
            ")": "(",
            "}": "{",
            "]": "[",
        }

        stack = []

        for bracket in parentheses:
            if bracket in bracket_pair:
                if not stack or stack.pop() != bracket_pair[bracket]:
                    return False
            elif bracket in "({[":
                stack.append(bracket)
            else:
                return False

        return len(stack) == 0


if __name__ == "__main__":
    inputs = {
        "case_1": ["()"],
        "case_2": ["(]"],
        "case_3": ["([)]"],
        "case_4": ["{[]}"],
    }

    outputs = {
        "case_1": True,
        "case_2": False,
        "case_3": False,
        "case_4": True,
    }

    solution = ValidateParentheses()

    for case, input in inputs.items():
        result = solution.is_valid_parentheses(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
