/* Valid Parentheses
Source		: https://leetcode.com/problems/valid-parentheses/
Level		: Easy
Description	: Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input
			string is valid. An input string is valid if:
				- Open brackets must be closed by the same type of brackets.
    			- Open brackets must be closed in the correct order.
    			- Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false
*/

package leetcode

var pairs = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func IsValid(s string) bool {
	opens := make([]string, 0)

	for _, p := range s {
		close := string(p)
		open, ok := pairs[close]
		if !ok {
			opens = append(opens, string(p))
			continue
		}

		if len(opens) == 0 || opens[len(opens)-1] != open {
			return false
		}

		opens[len(opens)-1] = ""
		opens = opens[:len(opens)-1]
	}

	return len(opens) == 0
}
