package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/runner"
)

var registry = map[string]runner.TestFunc{}

func main() {
	var solution = flag.String("solution", "", "Name of the solution to run (e.g., TwoSum)")
	flag.StringVar(solution, "s", *solution, "Alias for -solution")

	var list = flag.Bool("list", false, "List all available solutions")
	flag.BoolVar(list, "l", *list, "Alias for -list")

	flag.Parse()

	r := runner.NewSolutionRunner()
	r.RegisterSolutions(registry)
	// registerSolutions(&r)

	if *list {
		r.List()
		return
	}

	if *solution == "" {
		fmt.Println("Please specify a solution to run with -solution flag")
		fmt.Println("Example: go run . -solution TwoSum")
		fmt.Println("Use -list to see available solutions")
		os.Exit(1)
	}

	if err := r.Run(*solution); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func register(name string, fn runner.TestFunc) {
	if _, exists := registry[name]; exists {
		panic(fmt.Sprintf("solution %q is already registered", name))
	}
	registry[name] = fn
}
