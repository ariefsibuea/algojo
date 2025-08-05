"""
LeetCode Problem : Best Time to Buy and Sell Stock
Topic            : Array, Dynamic Programming
Level            : Easy
URL              : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
Description      : You are given an array prices where prices[i] is the price of a given stock on the ith day. You want
        to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to
        sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any
        profit, return 0.
Examples         :
        Example 1:
        Input: prices = [7,1,5,3,6,4]
        Output: 5
        Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

        Example 2:
        Input: prices = [7,6,4,3,1]
        Output: 0
        Explanation: In this case, no transactions are done and the max profit = 0.
"""

from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        """Finds the maximum profit by buying and selling a stock once at optimal days.

        Args:
            prices (List[int]): Array of stock prices where prices[i] is the price on day i.

        Returns:
            int: Maximum profit that can be achieved or 0 if no profit is possible.

        Time Complexity:
            O(n): Where n is the length of prices array, as we traverse the array once.

        Space Complexity:
            O(1): Only constant extra space is used for variables.
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
