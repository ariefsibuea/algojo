"""
LeetCode Problem : Longest Substring Without Repeating Characters
Topic            : Hash Table, String, Sliding Window
Level            : Medium
URL              : https://leetcode.com/problems/longest-substring-without-repeating-characters
Description      : Given a string s, find the length of the longest substring without duplicate characters.
Examples         :
                    Example 1:
                    Input: s = "abcabcbb"
                    Output: 3
                    Explanation: The answer is "abc", with the length of 3.

                    Example 2:
                    Input: s = "bbbbb"
                    Output: 1
                    Explanation: The answer is "b", with the length of 1.

                    Example 3:
                    Input: s = "pwwkew"
                    Output: 3
                    Explanation: The answer is "wke", with the length of 3.
"""


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """Finds the length of the longest substring without repeating characters using sliding window approach.

        Args:
            s (str): Input string to analyze.

        Returns:
            int: Length of the longest substring without repeating characters.

        Time Complexity:
            O(n): Where n is the length of the input string, as we only traverse the string once.

        Space Complexity:
            O(min(n, a)): Where a is the size of the character set.
        """
        char_map = {}
        max_length = 0
        start = 0

        for i, c in enumerate(s):
            if c in char_map and char_map[c] >= start:
                start = char_map[c] + 1

            char_map[c] = i
            max_length = max(max_length, (i - start) + 1)

        return max_length


def run_tests():
    inputs = {
        "case_1": ["abcabcbb"],
        "case_2": ["bbbbb"],
        "case_3": ["pwwkew"],
        "case_4": ["au"],
        "case_5": ["dvdf"],
    }
    outputs = {"case_1": 3, "case_2": 1, "case_3": 3, "case_4": 2, "case_5": 3}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.lengthOfLongestSubstring(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
