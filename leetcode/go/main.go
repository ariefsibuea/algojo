package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var solution = flag.String("solution", "", "Name of the solution to run (e.g., TwoSum)")
	var list = flag.Bool("list", false, "List all available solutions")

	flag.Parse()

	runner := NewSolutionRunner()

	if *list {
		runner.List()
		return
	}

	if *solution == "" {
		fmt.Println("Please specify a solution to run with --solution flag")
		fmt.Println("Example: go run . --solution TwoSum")
		fmt.Println("Use --list to see available solutions")
		os.Exit(1)
	}

	if err := runner.Run(*solution); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
