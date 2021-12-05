package validparentheses

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
