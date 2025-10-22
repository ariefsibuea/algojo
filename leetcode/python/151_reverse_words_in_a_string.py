"""
Problem          : Reverse Words in a String
Topics           : Two Pointers, String
Level            : Medium
URL              : https://leetcode.com/problems/reverse-words-in-a-string
Description      :
Examples         :
                    Example 1:
                    Input: s = "the sky is blue"
                    Output: "blue is sky the"

                    Example 2:
                    Input: s = "  hello world  "
                    Output: "world hello"
                    Explanation: Your reversed string should not contain leading or trailing spaces.

                    Example 3:
                    Input: s = "a good   example"
                    Output: "example good a"
                    Explanation: You need to reduce multiple spaces between two words to a single space in the
                    reversed string.
"""


class Solution:
    def reverseWords(self, s: str) -> str:
        words = s.split()

        left, right = 0, len(words) - 1
        while left < right:
            words[left], words[right] = words[right], words[left]
            left += 1
            right -= 1

        return " ".join(words)


def run_tests():
    inputs = {
        "case_1": ["the sky is blue"],
        "case_2": ["  hello world  "],
        "case_3": ["a good   example"],
    }
    outputs = {
        "case_1": "blue is sky the",
        "case_2": "world hello",
        "case_3": "example good a",
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.reverseWords(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
