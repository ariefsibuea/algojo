package main

/**
 * LeetCode Problem : Fruit Into Baskets
 * Topics           : Array, Hash Table, Sliding Window, Weekly Contest 102
 * Level            : Medium
 * URL              : https://leetcode.com/problems/fruit-into-baskets
 * Description      :
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
