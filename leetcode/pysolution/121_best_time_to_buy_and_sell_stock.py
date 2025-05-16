"""
LeetCode Problem : Best Time to Buy and Sell Stock
Topic            : Array, Dynamic Programming
Level            : Easy
URL              : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
"""

from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        """
        Given an array of prices representing stock prices where prices[i] is the price on day i,
        find the maximum profit you can achieve by choosing a single day to buy and a single day to sell.
        You must buy before selling. If no profit can be achieved, return 0.
        Args:
            prices (List[int]): Array of stock prices where prices[i] is the price on day i
        Returns:
            int: Maximum profit that can be achieved
        Example 1:
            Input: prices = [7,1,5,3,6,4]
            Output: 5
            Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
            Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
        Example 2:
            Input: prices = [7,6,4,3,1]
            Output: 0
            Explanation: In this case, no transactions are done and the max profit = 0.
        Solution:
            One-pass with tracking minimum
        Time Complexity:
            O(n) where n is the length of prices array
        Space Complexity:
            O(1) since we only use two variables
        """

        max_profit = 0
        min_price = float("inf")
        for price in prices:
            min_price = min(min_price, price)
            max_profit = max(max_profit, price - min_price)

        return max_profit


def run_tests():
    inputs = {"case_1": [[7, 1, 5, 3, 6, 4]], "case_2": [[7, 6, 4, 3, 1]]}
    outputs = {"case_1": 5, "case_2": 0}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.maxProfit(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
