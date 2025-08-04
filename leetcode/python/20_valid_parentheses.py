"""
LeetCode Problem : Valid Parentheses
Topic            : String, Stack
Level            : Easy
URL              : https://leetcode.com/problems/valid-parentheses/description
"""


class Solution:
    def isValid(self, s: str) -> bool:
        """Determines if the input string of parentheses is valid. A string is considered valid if:
            1. Open brackets are closed by the same type of brackets.
            2. Open brackets are closed in the correct order.
            3. Every close bracket has a corresponding open bracket of the same type.

        Args:
            s (str): The input string containing parentheses, braces, or brackets.

        Returns:
            bool: True if the input string is valid, False otherwise.

        Solution:
            Stack approach

        Time Complexity:
            O(n): Each character in the string is processed exactly once.

        Space Complexity:
            O(n): The stack can grow up to the size of the input string in the worst case.
        """

        bracket_pair = {
            ")": "(",
            "}": "{",
            "]": "[",
        }

        stack = []
        for c in s:
            if c in bracket_pair:
                if not stack or stack.pop() != bracket_pair[c]:
                    return False
            else:
                stack.append(c)

        return not stack


def run_tests():
    inputs = {"case_1": ["()"], "case_2": ["()[]{}"], "case_3": ["(]"], "case_4": ["([])"], "case_5": ["}"]}
    outputs = {"case_1": True, "case_2": True, "case_3": False, "case_4": True, "case_5": False}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.isValid(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
