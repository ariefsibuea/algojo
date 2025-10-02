package main

/**
 * LeetCode Problem : Maximum Points You Can Obtain From Cards
 * Topics           : Array, Sliding Window, Prefix Sum, Weekly Contest 186
 * Level            : Medium
 * URL              : https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/description/
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: cardPoints = [1,2,3,4,5,6,1], k = 3
 * 					Output: 12
 * 					Explanation: After the first step, your score will always be 1. However, choosing the rightmost
 * 					card first will maximize your total score. The optimal strategy is to take the three cards on the
 * 					right, giving a final score of 1 + 6 + 5 = 12.
 *
 * 					Example 2:
 * 					Input: cardPoints = [2,2,2], k = 2
 * 					Output: 4
 * 					Explanation: Regardless of which two cards you take, your score will always be 4.
 *
 * 					Example 3:
 * 					Input: cardPoints = [9,7,7,9,7,7,9], k = 7
 * 					Output: 55
 * 					Explanation: You have to take all the cards. Your score is the sum of points of all cards.
 */

func maxPointsFromCards(cardPoints []int, k int) int {
	totalPoints := 0
	for _, cardPoint := range cardPoints {
		totalPoints += cardPoint
	}

	if len(cardPoints) <= k {
		return totalPoints
	}

	deductPoints := 0
	maxPoints := 0
	numDeductCard := len(cardPoints) - k

	left := 0
	for right := 0; right < len(cardPoints); right++ {
		deductPoints += cardPoints[right]

		if right-left+1 == numDeductCard {
			maxPoints = max(maxPoints, totalPoints-deductPoints)
			deductPoints -= cardPoints[left]
			left += 1
		}
	}

	return maxPoints
}
