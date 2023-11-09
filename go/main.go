package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"reflect"

	"algojo.ariefsibuea.dev/leetcode"
	"algojo.ariefsibuea.dev/test/testcase"
)

type config struct {
	packageName string
	methodName  string
	testName    string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.packageName, "package", "", "Package name of solution")
	flag.StringVar(&cfg.methodName, "method", "", "Method name of solution")
	flag.StringVar(&cfg.testName, "test", "", "Test file name")
	flag.Parse()

	if cfg.packageName == "" {
		log.Fatalf("'packageName' is required")
	}
	if cfg.methodName == "" {
		log.Fatalf("'methodName' is required")
	}

	err := executeSolutionMethod(cfg.packageName, cfg.methodName, cfg.testName)
	if err != nil {
		log.Println(err)
	}
}

func executeSolutionMethod(packageName, methodName, testFileName string) error {
	var solution interface{}

	switch packageName {
	case "leetcode":
		solution = new(leetcode.Solution)
	default:
		return fmt.Errorf("package \"%s\"not found", packageName)
	}

	solutionMethod := reflect.ValueOf(solution).MethodByName(methodName)

	if testFileName == "" {
		log.Printf("no test file, set to run test")
		return nil
	}

	return executeTestCases(solutionMethod, packageName, testFileName)
}

func executeTestCases(solutionMethod reflect.Value, packageName, testFileName string) error {
	testCases := testcase.GetTestCases(packageName, testFileName)

	for _, testCase := range testCases {
		numTestParams := len(testCase.Inputs)
		numRequiredParams := solutionMethod.Type().NumIn()
		if numTestParams != numRequiredParams {
			return fmt.Errorf("number of test params: %d, required: %d", numTestParams, numRequiredParams)
		}

		inputs, err := convertValues(solutionMethod, testCase.Inputs)
		if err != nil {
			return fmt.Errorf("convert values of inputs: %v", err)
		}

		results := solutionMethod.Call(inputs)
		result := results[0].Interface()

		err = compareOutput(testCase.Name, result, testCase.Outputs[0])
		if err != nil {
			return fmt.Errorf("compare output: %v", err)
		}
	}

	return nil
}

func compareOutput(caseName string, output, expectedOutput interface{}) error {
	resultRaw, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("marshal output: %v", err)
	}

	expectedResultRaw, err := json.Marshal(expectedOutput)
	if err != nil {
		return fmt.Errorf("marshal expected output: %v", err)
	}

	if string(expectedResultRaw) != string(resultRaw) {
		fmt.Printf("==== Run test case %s, FAILED: expected output = %s, got %s\n", caseName, string(expectedResultRaw), string(resultRaw))
	} else {
		fmt.Printf("==== Run test case %s, SUCCESS\n", caseName)
	}

	return nil
}

func convertValues(method reflect.Value, inputs []interface{}) ([]reflect.Value, error) {
	convertedInputs := make([]reflect.Value, 0, len(inputs))
	for i := range inputs {
		val, err := convertValue(inputs[i], method.Type().In(i))
		if err != nil {
			return nil, err
		}

		convertedInputs = append(convertedInputs, val)
	}
	return convertedInputs, nil
}

func convertValue(input interface{}, targetType reflect.Type) (reflect.Value, error) {
	var converted reflect.Value

	inputVal := reflect.ValueOf(input)
	switch targetType.Kind() {
	case reflect.Slice:
		if inputVal.Kind() == reflect.Slice {
			sliceType := reflect.SliceOf(targetType.Elem())
			slice := reflect.MakeSlice(sliceType, inputVal.Len(), inputVal.Len())

			for i := 0; i < inputVal.Len(); i++ {
				var sliceItem reflect.Value
				if inputVal.Index(i).Elem().Type() != targetType.Elem() {
					sliceItem = reflect.New(targetType.Elem()).Elem()
					sliceItem.Set(inputVal.Index(i).Elem().Convert(targetType.Elem()))
				} else {
					sliceItem = inputVal.Index(i).Elem()
				}

				slice.Index(i).Set(sliceItem)
			}

			converted = slice
		} else {
			return reflect.Value{}, fmt.Errorf("cannot convert %v to %v", inputVal.Type(), targetType)
		}
	default:
		if targetType != inputVal.Type() {
			converted = reflect.New(targetType).Elem()
			converted.Set(inputVal.Convert(targetType))
		} else {
			converted = inputVal
		}
	}

	return converted, nil
}
