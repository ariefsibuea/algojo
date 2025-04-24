"""
LeetCode Problem : Longest Substring Without Repeating Characters
Topic            : Hash Table, String, Sliding Window
Level            : Medium
URL              : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
"""


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        Finds the length of the longest substring without repeating characters.
        Args:
            s (str): The input string.
        Returns:
            int: The length of the longest substring without repeating characters.
        Example:
            >>> solution.lengthOfLongestSubstring("abcabcbb")
            3
            >>> solution.lengthOfLongestSubstring("bbbbb")
            1
            >>> solution.lengthOfLongestSubstring("pwwkew")
            3
        Notes:
            - The function uses a sliding window approach with a hash map to track the last seen index of each character.
            - The time complexity is O(n), where n is the length of the input string.
            - The space complexity is O(min(n, a)), where a is the size of the character set.
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
