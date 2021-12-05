package addbinary

import (
	"strconv"
	"strings"
)

func AddBinary(a string, b string) string {
	ra := []rune(a)
	rb := []rune(b)

	indexA, indexB := len(ra)-1, len(rb)-1
	remainder, binA, binB := 0, 0, 0
	result := make([]string, 0)
	for indexA >= 0 || indexB >= 0 {
		binA, binB = 0, 0
		if indexA >= 0 {
			binA, _ = strconv.Atoi(string(ra[indexA]))
			indexA--
		}
		if indexB >= 0 {
			binB, _ = strconv.Atoi(string(rb[indexB]))
			indexB--
		}

		bin := (binA + binB + remainder) % 2
		remainder = (binA + binB + remainder) / 2
		result = append([]string{strconv.Itoa(bin)}, result...)
	}

	if remainder > 0 {
		result = append([]string{strconv.Itoa(remainder)}, result...)
	}

	return strings.Join(result, "")
}
