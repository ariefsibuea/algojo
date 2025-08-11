"""
LeetCode Problem : Top K Frequent Elements
Topic            : Array, Hash Table, Divide and Conquer, Sorting, Heap (Priority Queue), Bucket Sort, Counting, Quickselect
Level            : Medium
URL              : https://leetcode.com/problems/top-k-frequent-elements/
Description      : Given an integer array nums and an integer k, return the k most frequent elements. You may return
        the answer in any order.
Examples         :
        Example 1:
        Input: nums = [1,1,1,2,2,3], k = 2
        Output: [1,2]

        Example 2:
        Input: nums = [1], k = 1
        Output: [1]
"""

from collections import Counter
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        """Finds the k most frequent elements in the given list of integers using bucket sort.

        Args:
            nums (List[int]): A list of integers.
            k (int): The number of most frequent elements to return.

        Returns:
            List[int]: A list containing the k most frequent elements in descending order of frequency.

        Time Complexity:
            O(n): Where n is the length of the input list, as we iterate through the array once.

        Space Complexity:
            O(n): Space used for the hash map and frequency buckets.
        """

        count = Counter(nums)
        max_freq = max(count.values()) if count else 0
        buckets = [[] for _ in range(max_freq + 1)]

        for num, freq in count.items():
            buckets[freq].append(num)

        result = []
        for freq in range(max_freq, 0, -1):
            if buckets[freq]:
                result.extend(buckets[freq])
                if len(result) >= k:
                    break

        return result[:k]


def run_tests():
    inputs = {"case_1": [[1, 1, 1, 2, 2, 3], 2], "case_2": [[1], 1]}
    outputs = {"case_1": [1, 2], "case_2": [1]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.topKFrequent(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
