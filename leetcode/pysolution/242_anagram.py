"""
LeetCode Problem : Valid Anagram
Topic            : Hash Table, String, Sorting
Level            : Easy
URL              : https://leetcode.com/problems/valid-anagram/description/
"""

from typing import Any


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        """
        Determines if two strings are anagrams of each other.
        An anagram is a word or phrase formed by rearranging the letters of a
        different word or phrase, typically using all the original letters exactly once.
        Args:
            s (str): The first string to compare.
            t (str): The second string to compare.
        Returns:
            bool: True if the strings are anagrams, False otherwise.
        Example:
            isAnagram("anagram", "nagaram") -> True
            isAnagram("rat", "car") -> False
        Solution:
            Hash map approach
        Approach:
            - Check if the lengths of the two strings are different. If they are, return False.
            - Use a dictionary to count the occurrences of each character in the first string.
            - Iterate through the second string and decrement the count of each character in the dictionary.
            - If a character in the second string is not in the dictionary or its count becomes negative, return False.
            - Finally, check if the dictionary is empty, which indicates the strings are anagrams.
        Time Complexity: O(n)
        Space Complexity: O(n)
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
