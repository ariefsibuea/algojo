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
        """Finds the first bad version in a sequence of versions using binary search.
        
        Args:
            n (int): The total number of versions.
            
        Returns:
            int: The first bad version number.
            
        Time Complexity:
            O(log n): Binary search halves the search space in each iteration.
            
        Space Complexity:
            O(1): Only constant extra space is used.
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
