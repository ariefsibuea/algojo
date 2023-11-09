package testcase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type TestCase struct {
	Name    string
	Inputs  []interface{} `json:"input"`
	Outputs []interface{} `json:"output"`
}

func getFilePath(packageName, testFileName string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to get the current filename")
	}
	dirName := filepath.Dir(fileName)

	return fmt.Sprintf("%s/%s/%s.yaml", dirName, packageName, testFileName)
}

func GetTestCases(packageName, testFileName string) []TestCase {
	file, err := os.ReadFile(getFilePath(packageName, testFileName))
	if err != nil {
		log.Fatalf("read test file err: %v", err)
	}

	test := map[string]interface{}{}
	err = yaml.Unmarshal(file, &test)
	if err != nil {
		log.Fatalf("unmarshal test cases: %v", err)
	}

	mapTestCases, ok := test["TestCases"].(map[string]interface{})
	if !ok {
		log.Fatalf("unable to get TestCases data")
	}

	testCases := make([]TestCase, 0, len(mapTestCases))
	for n, c := range mapTestCases {
		rawCase, err := json.Marshal(c)
		if err != nil {
			log.Fatalf("marshal single test case: %v", err)
		}

		testCase := new(TestCase)
		err = json.Unmarshal(rawCase, &testCase)
		if err != nil {
			log.Fatalf("unmarshal single test case: %v", err)
		}

		testCase.Name = n
		testCases = append(testCases, *testCase)
	}

	return testCases
}
