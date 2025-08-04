"""
LeetCode Problem : Evaluate Reverse Polish Notation
Topic            : Array, Math, Stack
Level            : Medium
URL              : https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
"""

from typing import Any, List


class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        """Evaluates the value of an arithmetic expression in Reverse Polish Notation (RPN).
        
        Args:
            tokens (List[str]): An array of strings representing arithmetic expression in RPN.
            
        Returns:
            int: Result of evaluating the expression.
            
        Time Complexity:
            O(n): Where n is the number of tokens, as each token is processed once.
            
        Space Complexity:
            O(n): In the worst case, the stack may store up to n/2 numbers.
        """

        number_stack = []
        valid_operators = {
            "+": lambda a, b: a + b,
            "-": lambda a, b: a - b,
            "*": lambda a, b: a * b,
            "/": lambda a, b: int(a / b),
        }

        for token in tokens:
            if token not in valid_operators:
                number_stack.append(int(token))
            else:
                b = number_stack.pop()
                a = number_stack.pop()
                number_stack.append(valid_operators[token](a, b))

        return number_stack[0]


def run_tests():
    inputs = {
        "case_1": [["2", "1", "+", "3", "*"]],
        "case_2": [["4", "13", "5", "/", "+"]],
        "case_3": [["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]],
    }
    outputs = {"case_1": 9, "case_2": 6, "case_3": 22}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.evalRPN(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
