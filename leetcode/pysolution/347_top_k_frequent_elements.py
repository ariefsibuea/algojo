"""
LeetCode Problem : Top K Frequent Elements
Topic            : Array, Hash Table, Divide and Conquer, Sorting, Heap (Priority Queue), Bucket Sort, Counting, Quickselect
Level            : Medium
URL              : https://leetcode.com/problems/top-k-frequent-elements/description/
"""

from collections import Counter
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        """
        Finds the k most frequent elements in the given list of integers.

        Args:
            nums (List[int]): A list of integers.
            k (int): The number of most frequent elements to return.

        Returns:
            List[int]: A list containing the k most frequent elements in descending order of frequency.
            If there are ties in frequency, the order of elements in the result may vary.

        Example:
            >>> topKFrequent([1, 1, 1, 2, 2, 3], 2)
            [1, 2]

        Notes:
            - The function uses a bucket sort approach to group elements by their frequency.
            - The time complexity is O(n), where n is the length of the input list, assuming the number of unique elements is much smaller than n.
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
