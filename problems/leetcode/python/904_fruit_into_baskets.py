"""
LeetCode Problem : Fruit Into Baskets
Topics           : Array, Hash Table, Sliding Window
Level            : Medium
URL              : https://leetcode.com/problems/fruit-into-baskets
Description      : You are given an integer array fruits where fruits[i] is the type of fruit the i-th tree produces.
                    You want to collect as much fruit as possible. However, the owner has some strict rules that you
                    must follow:
                        - You only have two baskets, and each basket can only hold a single type of fruit.
                        - Starting from any tree of your choice, you must pick exactly one fruit from every tree
                            (including the start tree) while moving to the right.
                        - The picked fruits must fit in one of your baskets.
                        - Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
                    Given the integer array fruits, return the maximum number of fruits you can pick.
Examples         :
                    Example 1:
                    Input: fruits = [1,2,1]
                    Output: 3
                    Explanation: We can pick from all 3 trees.

                    Example 2:
                    Input: fruits = [0,1,2,2]
                    Output: 3
                    Explanation: We can pick from trees [1,2,2].
                    If we had started at the first tree, we would only pick from trees [0,1].

                    Example 3:
                    Input: fruits = [1,2,3,2,2]
                    Output: 4
                    Explanation: We can pick from trees [2,3,2,2].
                    If we had started at the first tree, we would only pick from trees [1,2].
"""

from typing import List


class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        left, right = 0, 0
        max_fruit = 0
        baskets: dict[int, int] = {}

        while right <= len(fruits) - 1:
            baskets[fruits[right]] = baskets.get(fruits[right], 0) + 1

            while len(baskets) > 2:
                baskets[fruits[left]] -= 1
                if baskets[fruits[left]] == 0:
                    del baskets[fruits[left]]
                left += 1

            max_fruit = max(max_fruit, right - left + 1)
            right += 1

        return max_fruit


def run_tests():
    inputs = {
        "case_1": [
            [1, 2, 1],
        ],
        "case_2": [
            [0, 1, 2, 2],
        ],
        "case_3": [
            [1, 2, 3, 2, 2],
        ],
    }
    outputs = {
        "case_1": 3,
        "case_2": 3,
        "case_3": 4,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.totalFruit(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
