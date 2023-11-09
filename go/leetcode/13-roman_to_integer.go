/* Roman to Integer
Source		: https://leetcode.com/problems/roman-to-integer/
Level		: Easy
Description	: Given a roman numeral, convert it to an integer.

Example 1:
Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

package leetcode

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func (soln Solution) RomanToInt(s string) int {
	result, lastnum, currnum := 0, 0, 0
	for _, v := range s {
		currnum = romans[string(v)]
		if currnum > lastnum {
			result -= lastnum
			result = result + (currnum - lastnum)
		} else {
			result = result + currnum
		}

		lastnum = currnum
	}

	return result
}
