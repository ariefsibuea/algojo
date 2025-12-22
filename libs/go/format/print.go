package format

import (
	"fmt"
	"strings"
)

func PrintInput(inputs map[string]interface{}) {
	strInputs := []string{}
	for k, v := range inputs {
		strInputs = append(strInputs, fmt.Sprintf("%s = %+v", k, v))
	}
	strInput := strings.Join(strInputs, ", ")

	fmt.Printf("=== \033[1mINPUT:\033[0m %s\n", strInput)
}

func PrintFailed(format string, args ...interface{}) {
	fmt.Printf("=== \033[1;31mFAILED:\033[0m %s\n", fmt.Sprintf(format, args...))
}

func PrintSuccess(format string, args ...interface{}) {
	fmt.Printf("=== \033[1;32mSUCCESS:\033[0m %s\n", fmt.Sprintf(format, args...))
}
