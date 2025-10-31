"""
Problem          : Two Sum Less Than K
Topics           : Sort, Search
Level            : Easy
URL              : https://leetcode.com/problems/two-sum-less-than-k
Description      : Given an array of integers, nums, and an integer k, find the maximum sum of two elements in nums
                    less than k. Otherwise, return −1 if no such pair exists.
Examples         :
                    Example 1:
                    Input: nums = [2,1,3,3,5], k = 4
                    Output: 3
                    Explanation: The maximum sum less than k is 3, i.e., 2+1 = 3.

                    Example 2:
                    Input: nums = [8,4,9,2,10,1], k = 9
                    Output: 6
                    Explanation: The maximum sum less than k is 6, i.e., 4+2 = 6.

                    Example 3:
                    Input: nums = [3,4,5,8,6,2], k = 5
                    Output: -1

                    Example 4:
                    Input: nums = [4,4,4,4,4,4,4,4], k = 12
                    Output: 8
"""


class Solution:
    def two_sum_less_than_k(self, nums: list[int], k: int) -> int:
        sorted(nums)
        if nums[0] >= k:
            return -1

        nums_len = len(nums)
        result = -1

        for i, num in enumerate(nums):
            j, want = -1, k - num
            left, right = i + 1, nums_len - 1
            while left <= right:
                mid = (left + right) // 2
                if nums[mid] < want:
                    j = mid
                    left = mid + 1
                else:
                    right = mid - 1

            if j > i:
                result = max(result, num + nums[j])

        return result


def run_tests():
    inputs = {
        "case-1": [
            [2, 1, 3, 3, 5],
            4,
        ],
        "case-2": [
            [8, 4, 9, 2, 10, 1],
            9,
        ],
        "case-3": [
            [3, 4, 5, 8, 6, 2],
            5,
        ],
        "case-4": [
            [4, 4, 4, 4, 4, 4, 4, 4],
            12,
        ],
    }
    outputs = {
        "case-1": 3,
        "case-2": 6,
        "case-3": -1,
        "case-4": 8,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.two_sum_less_than_k(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
