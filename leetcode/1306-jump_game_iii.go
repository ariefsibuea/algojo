/* Jump Game III
Source		: https://leetcode.com/problems/jump-game-iii/
Level		: Medium
Description	:
	Given an array of non-negative integers arr, you are initially positioned at start index of the array. When you are
	at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach to any index with value 0.
	Notice that you can not jump outside of the array at any time.

Example 1:
Input: arr = [4,2,3,0,3,1,2], start = 5
Output: true
Explanation:
All possible ways to reach at index 3 with value 0 are:
index 5 -> index 4 -> index 1 -> index 3
index 5 -> index 6 -> index 4 -> index 1 -> index 3

Example 2:
Input: arr = [4,2,3,0,3,1,2], start = 0
Output: true
Explanation:
One possible way to reach at index 3 with value 0 is:
index 0 -> index 4 -> index 1 -> index 3

Example 3:
Input: arr = [3,0,2,1,2], start = 2
Output: false
Explanation: There is no way to reach at index 1 with value 0.
*/

package leetcode

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(e int) {
	q.Elements = append(q.Elements, e)
}

func (q *Queue) Len() int {
	return len(q.Elements)
}

func (q *Queue) Dequeue() int {
	e := q.Elements[0]
	q.Elements[0] = 0
	q.Elements = q.Elements[1:]
	return e
}

func CanReach(arr []int, start int) bool {
	minIndex, maxIndex := 0, len(arr)-1

	q := Queue{}
	isMarked := make(map[int]bool)

	q.Enqueue(start)
	isMarked[start] = true

	for q.Len() > 0 {
		e := q.Dequeue()
		if arr[e] == 0 {
			return true
		}

		index := e + arr[e]
		if index >= minIndex && index <= maxIndex && !isMarked[index] {
			q.Enqueue(index)
			isMarked[index] = true
		}

		index = e - arr[e]
		if index >= minIndex && index <= maxIndex && !isMarked[index] {
			q.Enqueue(index)
			isMarked[index] = true
		}
	}

	return false
}
