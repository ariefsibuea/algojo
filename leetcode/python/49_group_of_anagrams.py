"""
LeetCode Problem : Group Anagrams
Topic            : Array, Hash Table, String, Sorting
Level            : Medium
URL              : https://leetcode.com/problems/group-anagrams/
"""

from collections import defaultdict
from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        """Groups strings that are anagrams of each other from a given list.

        Args:
            strs (List[str]): A list of strings to be grouped.

        Returns:
            List[List[str]]: A list of lists where each inner list contains strings that are anagrams of each other.

        Time Complexity:
            O(N * K): Where N is the length of input array and K is the maximum length of a string in the input.

        Space Complexity:
            O(N * K): Space used to store frequency tuples and grouped anagrams.
        """

        # defaultdict is a subclass of the built-in dict class from the collection module. It is used to provide a default
        # value for a nonexistnt key in the dictionary, eliminating the need for checking if the key exists before using
        # it.
        anagram_map = defaultdict(list)
        for s in strs:
            s_count = [0] * 26
            for char in s:
                s_count[ord(char) - ord("a")] += 1
            key = tuple(s_count)
            anagram_map[key].append(s)
        return list(anagram_map.values())


def run_tests():
    inputs = {"case_1": [["eat", "tea", "tan", "ate", "nat", "bat"]], "case_2": [[""]]}
    outputs = {"case_1": [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]], "case_2": [[""]]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.groupAnagrams(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
