"""
LeetCode Problem : Longest Repeating Character Replacement
Topics           : Hash Table, String, Sliding Window
Level            : Medium
URL              : https://leetcode.com/problems/longest-repeating-character-replacement
Description      : You are given a string s and an integer k. You can choose any character of the string and change it
                    to any other uppercase English character. You can perform this operation at most k times. Return
                    the length of the longest substring containing the same letter you can get after performing the
                    above operations.
Examples         :
                    Example 1:
                    Input: s = "ABAB", k = 2
                    Output: 4
                    Explanation: Replace the two 'A's with two 'B's or vice versa.

                    Example 2:
                    Input: s = "AABABBA", k = 1
                    Output: 4
                    Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
                    The substring "BBBB" has the longest repeating letters, which is 4.
                    There may exists other ways to achieve this answer too.
Reference	     : https://www.hellointerview.com/learn/code/sliding-window/longest-repeating-character-replacement
"""


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        char_count = {}
        max_freq = 0
        max_len = 0

        left = 0
        for right, char in enumerate(s):
            char_count[char] = char_count.get(char, 0) + 1
            max_freq = max(max_freq, char_count[char])

            while (right - left + 1) - max_freq > k:
                left_char = s[left]
                char_count[left_char] -= 1
                left += 1

            max_len = max(max_len, right - left + 1)

        return max_len


def run_tests():
    inputs = {
        "case_1": [
            "ABAB",
            2,
        ],
        "case_2": [
            "AABABBA",
            1,
        ],
    }
    outputs = {
        "case_1": 4,
        "case_2": 4,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.characterReplacement(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
