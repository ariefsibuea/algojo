package romantointeger

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
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
