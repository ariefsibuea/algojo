package lengthoflastworld

func LengthOfLastWord(s string) int {
	r := []rune(s)
	length, index := 0, len(r)-1

	for index >= 0 && string(r[index]) == " " {
		index--
	}

	for index >= 0 && string(r[index]) != " " {
		length++
		index--
	}

	return length
}
