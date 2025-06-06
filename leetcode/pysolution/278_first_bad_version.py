"""
LeetCode Problem : First Bad Version
Topic            : Binary Search, Interactive
Level            : Easy
URL              : https://leetcode.com/problems/first-bad-version/description
"""


# The isBadVersion API is already defined for you.
def isBadVersion(version: int) -> bool:
    return version >= 4


class Solution:
    def firstBadVersion(self, n: int) -> int:
        """
        Find first version where code broken by using binary search approach.
        Assuming the version order goes from 1 to n.
        Args:
            n: int representing the total number of versions
        Returns:
            int: the first bad version number
        Examples:
            n = 5, bad = 4 -> output = 4
            All versions from 4 onwards are bad
        Solutions:
            Binary Search
        Time Complexity:
            O(log n) - binary search halves search space each iteration
        Space Complexity:
            O(1) - only uses a constant amount of extra space
        """

        low, high = 1, n

        while low < high:
            mid = low + (high - low) // 2

            if isBadVersion(mid):
                high = mid
            else:
                low = mid + 1

        return low


def run_tests():
    inputs = {"case_1": [5]}
    outputs = {"case_1": 4}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.firstBadVersion(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
