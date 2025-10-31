package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Fruit Into Baskets
 * Topics           : Array, Hash Table, Sliding Window
 * Level            : Medium
 * URL              : https://leetcode.com/problems/fruit-into-baskets
 * Description      : You are given an integer array fruits where fruits[i] is the type of fruit the i-th tree
 * 					produces. You want to collect as much fruit as possible. However, the owner has some strict rules
 * 					that you must follow:
 *                    - You only have two baskets, and each basket can only hold a single type of fruit.
 *                    - Starting from any tree of your choice, you must pick exactly one fruit from every tree
 * 						(including the start tree) while moving to the right.
 *                    - The picked fruits must fit in one of your baskets.
 *                    - Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
 *                  Given the integer array fruits, return the maximum number of fruits you can pick.
 * Examples         :
 * 					Example 1:
 * 					Input: fruits = [1,2,1]
 * 					Output: 3
 * 					Explanation: We can pick from all 3 trees.
 *
 * 					Example 2:
 * 					Input: fruits = [0,1,2,2]
 * 					Output: 3
 * 					Explanation: We can pick from trees [1,2,2].
 * 					If we had started at the first tree, we would only pick from trees [0,1].
 *
 * 					Example 3:
 * 					Input: fruits = [1,2,3,2,2]
 * 					Output: 4
 * 					Explanation: We can pick from trees [2,3,2,2].
 * 					If we had started at the first tree, we would only pick from trees [1,2].
 */

func totalFruit(fruits []int) int {
	// NOTE:
	// 	1. The time complexity of this function is O(n). Eventhough there is a nested loop, it at most access the array
	// 		n times.
	// 	2. The space complexity is O(1). There will be no more than 3 keys of the map.
	baskets := map[int]int{}
	maxNFruit := 0
	left := 0

	for right, fruit := range fruits {
		baskets[fruit] += 1

		for len(baskets) > 2 {
			prevFruit := fruits[left]
			baskets[prevFruit] -= 1
			if baskets[prevFruit] == 0 {
				delete(baskets, prevFruit)
			}
			left += 1
		}

		maxNFruit = max(maxNFruit, right-left+1)
	}

	return maxNFruit
}

func RunTestTotalFruit() {
	testCases := map[string]struct {
		fruits []int
		expect int
	}{
		"case-1": {
			fruits: []int{1, 2, 1},
			expect: 3,
		},
		"case-2": {
			fruits: []int{0, 1, 2, 2},
			expect: 3,
		},
		"case-3": {
			fruits: []int{1, 2, 3, 2, 2},
			expect: 4,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := totalFruit(testCase.fruits)
		if !cmp.EqualNumbers(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}

	fmt.Printf("\n✅ All tests passed!\n")
}
