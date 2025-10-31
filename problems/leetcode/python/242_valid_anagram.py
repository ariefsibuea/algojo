"""
LeetCode Problem : Valid Anagram
Topic            : Hash Table, String, Sorting
Level            : Easy
URL              : https://leetcode.com/problems/valid-anagram/
Description      : Given two strings s and t, return true if t is an anagram of s, and false otherwise. An anagram is
        a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the
        original letters exactly once.
Examples         :
        Example 1:
        Input: s = "anagram", t = "nagaram"
        Output: true

        Example 2:
        Input: s = "rat", t = "car"
        Output: false
"""

from typing import Any


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        """Determines if two strings are anagrams of each other.

        Args:
            s (str): The first string to compare.
            t (str): The second string to compare.

        Returns:
            bool: True if the strings are anagrams, False otherwise.

        Time Complexity:
            O(n): Where n is the length of the strings, as we iterate through each character.

        Space Complexity:
            O(1): Fixed-size space for character count (at most 26 lowercase letters).
        """

        if len(s) != len(t):
            return False

        char_count = {}
        for c in s:
            char_count[c] = char_count.get(c, 0) + 1

        for c in t:
            if char_count.get(c, 0) == 0:
                return False
            char_count[c] -= 1
            if char_count[c] == 0:
                char_count.pop(c, None)

        return True


def run_tests():
    inputs = {"case_1": ["anagram", "nagaram"], "case_2": ["rat", "car"], "case_3": ["ab", "a"]}
    outputs = {"case_1": True, "case_2": False, "case_3": False}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.isAnagram(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
