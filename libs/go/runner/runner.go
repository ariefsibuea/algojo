package runner

import (
	"fmt"
	"strings"
)

type TestFunc func()

type SolutionRunner struct {
	tests map[string]TestFunc
}

func NewSolutionRunner() SolutionRunner {
	return SolutionRunner{
		tests: map[string]TestFunc{},
	}
}

func (r SolutionRunner) Run(solutionName string) error {
	testFunc, exists := r.tests[solutionName]
	if !exists {
		return fmt.Errorf("solution '%s' does not exist", solutionName)
	}

	fmt.Printf("Running solution: %s\n", solutionName)
	fmt.Println(strings.Repeat("=", 50))
	testFunc()
	fmt.Println()

	return nil
}

func (r SolutionRunner) List() {
	fmt.Println("Available solutions:")
	for name := range r.tests {
		fmt.Printf("  - %s\n", name)
	}
}

func (r *SolutionRunner) RegisterSolutions(solutions map[string]TestFunc) {
	r.tests = solutions
}

func (r *SolutionRunner) AddSolution(name string, fn TestFunc) {
	r.tests[name] = fn
}
