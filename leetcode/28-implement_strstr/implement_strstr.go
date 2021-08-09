package implementstrstr

import "strings"

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	startIndex := 0
	result := -1
	for {
		if len(haystack) <= len(needle) && haystack != needle {
			result = -1
			break
		}

		if strings.HasPrefix(haystack, needle) {
			result = startIndex
			break
		}

		haystack = haystack[1:]
		startIndex = startIndex + 1
	}

	return result
}
