package permutationinstring

/**
 * Problem source: https://leetcode.com/problems/permutation-in-string/
 * Agenda:
 * 		Study Plan: Algorithm I
 *		Day I Sliding Window
 *		Level: Medium
 * Solution source:
**/

// CheckInclusion implements sliding window technique to solve check permutation of s1 in s2 problem.
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	charOfS1 := make([]int, 26) // based on length of a-z
	for i := range s1 {
		charOfS1[int(s1[i]-'a')] += 1
	}

	permutation := make([]int, 26) // based on length of a-z
	for i := range s1 {
		permutation[int(s2[i]-'a')] += 1
	}
	if isEqualArray(charOfS1, permutation) {
		return true
	}

	lenS1 := len(s1)
	for i := lenS1; i < len(s2); i++ {
		permutation[int(s2[i-lenS1]-'a')] -= 1
		permutation[int(s2[i]-'a')] += 1
		if isEqualArray(charOfS1, permutation) {
			return true
		}
	}
	return false
}

func isEqualArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
