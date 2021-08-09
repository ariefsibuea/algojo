package longestcommonprefix

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(prefix) > len(strs[i]) {
			prefix = prefix[0:len(strs[i])]
		}

		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
		}
	}

	return prefix
}
