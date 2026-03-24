package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
	"github.com/ariefsibuea/algojo/libs/go/format"
	"github.com/ariefsibuea/algojo/libs/go/runner"
)

func init() {
	register("StringLed", RunTestStringLed)
}

/*
 * Problem	: String LED
 * Topics	: Graph, Union-Find, Disjoint Set Union (DSU), Connectivity
 * Level	: 1
 * URL		: https://judge.beecrowd.com/en/problems/view/3171
 *
 * Description:
 * 		Mariazinha wants to arrange her Christmas tree with LED strings that were cut into several pieces by her
 * 		younger sister. Given N LED string segments (labeled 1 to N) and L connections between segments, determine if
 * 		all segments are connected together to form a complete LED string, or if any segment is missing/disconnected.
 *
 * Constraints:
 * 		- 1 ≤ N ≤ 100 (number of LED string segments)
 * 		- 1 ≤ L ≤ 100 (number of connections)
 * 		- Segments are numbered from 1 to N
 * 		- Each connection is between two different segments
 *
 * Examples:
 * 		Input:
 * 			4 3
 * 			1 2
 * 			3 4
 * 			1 3
 * 		Output: COMPLETO
 *
 * 		Input:
 * 			4 2
 * 			1 2
 * 			2 3
 * 		Output: INCOMPLETO
 */

func execStringLed() string {
	return execStringLedDFS()
}

func execStringLedDFS() string {
	var n, l int
	fmt.Scanf("%d %d", &n, &l)

	adjacent := make([][]int, n+1)
	for i := 0; i < l; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		adjacent[x] = append(adjacent[x], y)
		adjacent[y] = append(adjacent[y], x)
	}

	visited := make([]bool, n+1)
	stack := []int{1}

	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[i] {
			continue
		}

		visited[i] = true

		for _, neighbor := range adjacent[i] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	count := 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			count++
		}
	}

	if count != n {
		return "INCOMPLETO"
	}
	return "COMPLETO"
}

func execStringLedUnionFind() string {
	var n, l int
	fmt.Scanf("%d %d", &n, &l)

	roots := make([]int, n+1)
	for i := 1; i <= n; i++ {
		roots[i] = i
	}

	var find func(roots []int, i int) int
	find = func(roots []int, i int) int {
		if roots[i] != i {
			roots[i] = find(roots, roots[i])
		}
		return roots[i]
	}

	union := func(roots []int, x, y int) {
		rootX := find(roots, x)
		rootY := find(roots, y)
		if rootX != rootY {
			roots[rootY] = rootX
		}
	}

	for i := 0; i < l; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		union(roots, x, y)
	}

	root := find(roots, 1)
	for i := 2; i <= n; i++ {
		if find(roots, i) != root {
			return "INCOMPLETO"
		}
	}

	return "COMPLETO"
}

func RunTestStringLed() {
	runner.InitMetrics("StringLed")

	testCases := map[string]struct {
		input  []byte
		expect string
	}{
		"completo": {
			input:  []byte("4 3\n1 2\n3 4\n1 3"),
			expect: "COMPLETO",
		},
		"incompleto": {
			input:  []byte("4 2\n1 2\n2 3"),
			expect: "INCOMPLETO",
		},
		"single segment": {
			input:  []byte("1 0"),
			expect: "COMPLETO",
		},
		"all connected linear": {
			input:  []byte("5 4\n1 2\n2 3\n3 4\n4 5"),
			expect: "COMPLETO",
		},
	}

	var passedCount uint16 = 0

	for name, tc := range testCases {
		fmt.Printf("RUN %s\n", name)
		format.PrintInput(map[string]any{"input": string(tc.input)})

		r, w, _ := os.Pipe()
		oldStdin := os.Stdin
		os.Stdin = r
		w.Write(tc.input)
		w.Close()

		result := runner.ExecCountMetrics(execStringLed).(string)
		os.Stdin = oldStdin

		if !cmp.EqualStrings(result, tc.expect) {
			format.PrintFailed("expect = %v - got = %v\n", tc.expect, result)
			continue
		}

		format.PrintSuccess("test case '%s' passed", name)
		passedCount++
	}

	fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
	runner.PrintMetrics()
}
