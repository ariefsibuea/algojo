# Go Solution Specification

This document defines the structure, conventions, and coding standards for Go solutions in the `algojo` coding challenge repository.

## 1. Directory Structure

Solutions are organized by platform and language:

```text
problems/
├── leetcode/go/       # LeetCode solutions
├── hackerrank/go/     # HackerRank solutions
├── etc/go/            # Solutions from other sources (Interviews, text books)
├── docs/              # Documentation
├── template.go        # Solution template
└── template.py        # Python template (reference only)
```

## 2. Naming Conventions

Strict adherence to these conventions is required for the runner to correctly identify and execute your solutions.

| Element                  | Convention                                                       | Example                                    |
| ------------------------ | ---------------------------------------------------------------- | ------------------------------------------ |
| **File Name**            | `snake_case` (optional number prefix)                            | `two_sum.go`, `1_two_sum.go`               |
| **Solution Function**    | `camelCase`                                                      | `subarraySum`, `isBalancedBrackets`        |
| **Test Function**        | `RunTest` + `PascalCaseProblemName`                              | `RunTestSubarraySumEqualsK`                |
| **Constants**            | `PascalCase` (exported) or `camelCase` (unexported)              | `Valid4DigitPin`, `maxRetries`             |
| **Types/Structs**        | `PascalCase`                                                     | `LRUCache`, `ListNode`                     |
| **Method Receivers**     | 1-2 letter abbreviation                                          | `func (l *LRUCache)`, `func (n *ListNode)` |
| **Multi-implementation** | `functionName` + `_` + `PascalCaseApproach`                      | `findDuplicate_TortoiseHare`               |
| **Test Case Names**      | `kebab-case`                                                     | `"case-1"`, `"valid-4-digit"`              |
| **Test Input Helper**    | `build` + `PascalCaseProblem` + `PascalCaseTestName` + `[Field]` | `buildSplitListEvenNodesHead`              |

### Initialisms

Keep initialisms consistent in casing:

| Context    | Correct             | Incorrect           |
| ---------- | ------------------- | ------------------- |
| Exported   | `URL`, `HTTP`, `ID` | `Url`, `Http`, `Id` |
| Unexported | `url`, `http`, `id` | `uRL`, `hTTP`, `iD` |

**Examples:**

- `func getURL()` ✓ vs `func getUrl()` ✗
- `userID` ✓ vs `userId` ✗

## 3. File Structure

Every solution file must follow this exact order:

1.  **Package Declaration** (`package main`)
2.  **Imports** (Standard -> Project -> Side-effects)
3.  **Doc Comment** (Metadata block)
4.  **Constants** (Optional)
5.  **Types/Structs** (Optional)
6.  **Solution Function(s)**
7.  **Test Function** (`RunTestXxx`)
8.  **Helper Functions** (Optional)

---

## 4. Component Details

### 4.1 Package Declaration & Imports

All solution files belong to `package main`. Imports must be grouped and ordered:

1.  **Standard Library**
2.  **Project Packages**
3.  **Side-effect Imports**

Separate groups with a blank line.

```go
package main

import (
    "fmt"
    "os"
    "sort"

    "github.com/ariefsibuea/algojo/libs/go/cmp"
    "github.com/ariefsibuea/algojo/libs/go/format"
    "github.com/ariefsibuea/algojo/libs/go/runner"
)
```

### 4.2 Doc Comment

Every file must start with a standardized documentation block. This metadata is parsed by tools and humans alike.

#### Formatting Rules

1.  **Alignment**: Use tabs (`\t`) to align content after the colon.
2.  **Section Headers**: Use `Description:`, `Constraints:`, and `Examples:` as separate sections.
3.  **Continuation**: For multi-line fields (Description, Examples), start new lines with ` *\t` to align with the text.
4.  **Max Width**: Try to keep lines under 120 characters.
5.  **Separators**: Use ` *` (without trailing text) to separate logical sections within a field (e.g., between Example 1 and Example 2).

#### Example

```go
/*
 * Problem	: Task Scheduler
 * Topics	: Array, Hash Table, Greedy, Sorting, Priority Queue (Heap), Counting
 * Level	: Medium
 * URL		: https://leetcode.com/problems/task-scheduler/
 *
 * Description:
 * 		Schedule CPU tasks (labeled A-Z) with a cooldown constraint. Given a list of tasks and a cooldown period n, find
 * 		the minimum time needed to execute all tasks. Each time unit can either process one task or be idle. The same
 * 		task cannot be executed again until n time units have passed since its last execution. Tasks can be performed
 * 		in any order, and idle periods may be necessary to satisfy the cooldown constraint. Return the total number of
 * 		time units required.
 *
 * Constraints:
 * 		- 1 <= tasks.length <= 10^4
 * 		- tasks[i] is an uppercase English letter
 * 		- 0 <= n <= 100
 *
 * Examples:
 * 		Example 1:
 * 		Input: tasks = ["A","A","A","B","B","B"], n = 2
 * 		Output: 8
 * 		Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
 * 		After completing task A, you must wait two intervals before doing A again.
 *
 * 		Example 2:
 * 		Input: tasks = ["A","C","A","B","D","B"], n = 1
 * 		Output: 6
 * 		Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
 */
```

### 4.3 Constants & Types

Define local types and constants here. Shared types (like `ListNode`) are in `structures.go` and do not need re-definition.

```go
const (
    ValidPin      = 0
    InvalidLength = 1
)

type LRUCache struct {
    capacity int
    cache    map[int]*DLLNode
}
```

### 4.4 Solution Function

The primary logic. Use clear argument names.

If you implement **multiple approaches** for the same problem, use the suffix naming convention: `_` + `PascalCaseApproach`.

```go
// Primary/Best approach
func leastInterval(tasks []byte, n int) int {
    return leastInterval_PriorityQueue(tasks, n)
}

// Specific Approach 1
func leastInterval_PriorityQueue(tasks []byte, n int) int {
    // ...
}

// Specific Approach 2
func leastInterval_FillTheSlotAndSort(tasks []byte, n int) int {
    // ...
}
```

### 4.5 Test Function

The entry point for the runner. Must start with `RunTest` followed by the `PascalCase` problem name.

**Key Requirements:**

1.  **Metric Initialization:** Call `runner.InitMetrics("Title")`.
2.  **Test Cases:** Use a map of anonymous structs.
3.  **Execution:** Iterate through the map.
4.  **Assertions:** Use `cmp` package.
5.  **Output:** Use `format` package.
6.  **Failure:** Use `continue` (do not `os.Exit`).

```go
func RunTestTaskScheduler() {
    runner.InitMetrics("TaskScheduler")

    testCases := map[string]struct {
        tasks  []byte
        n      int
        expect int
    }{
        "example-1-basic": {
            tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
            n:      2,
            expect: 8,
        },
    }

    var passedCount int
    for name, tc := range testCases {
        fmt.Printf("RUN %s\n", name)
        format.PrintInput(map[string]interface{}{"tasks": string(tc.tasks), "n": tc.n})

        // Use runner.ExecCountMetrics to auto-measure memory/time
        result := runner.ExecCountMetrics(leastInterval, tc.tasks, tc.n).(int)
        if !cmp.EqualNumbers(result, tc.expect) {
            format.PrintFailed("expect = %v - got = %v", tc.expect, result)
            continue
        }

        format.PrintSuccess("test case '%s' passed", name)
        passedCount++
    }

    fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
    runner.PrintMetrics()
}
```

### 4.6 Input Helper Functions

Define a helper function **only** when the input requires complex initialization logic (e.g., circular linked lists, graphs with cycles, specific tree structures) that would clutter the test case definition.

**Do:**

- Use for complex types (`*ListNode` with cycles, `*TreeNode` with specific pointers).
- Use the naming convention `build` + `ProblemName` + `TestName` + `[Field]`.

**Don't:**

- Create helpers for simple types (arrays, slices, maps, primitives). Define these directly in the struct literal.

```go
// GOOD: Complex input construction
func buildSplitACircularLinkedListEvenNodesHead() *ListNode {
    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node1.Next = node2
    node2.Next = node1 // Cycle
    return node1
}

// BAD: Unnecessary helper for simple input
func buildTwoSumCase1Nums() []int {
    return []int{2, 7, 11, 15}
}
```

---

## 5. Solution Registration

After creating your file, you must register it in the `main.go` file of the corresponding directory (e.g., `problems/leetcode/go/main.go`).

1.  Open `main.go`.
2.  Locate the `registerSolutions` function.
3.  Add your `RunTestXxx` function to the map.
4.  **Sort alphabetically** by the key string.

```go
func registerSolutions(r *runner.SolutionRunner) {
    solutions := map[string]runner.TestFunc{
        // ... existing solutions ...
        "MinimumSwaps2":        RunTestMinimumSwaps2,
        "NewYearChaos":         RunTestNewYearChaos,
        "TaskScheduler":        RunTestTaskScheduler,  // <--- Your new solution
        "TwoStrings":           RunTestTwoStrings,
    }
    r.RegisterSolutions(solutions)
}
```

---

## 6. Coding Style

### Variable Names

Variable name length should be proportional to its scope size.

- **Small scope (1-7 lines):** `i` (index), `n` (node/size), `x` (value). Short names are acceptable.
- **Medium scope (8-15 lines):** `count`, `result`, `sum`, `task`. Single words that describe purpose.
- **Large scope (15+ lines):** `userCount`, `processedItems`, `remainingTasks`. Descriptive names are required.

### Variable Declaration

- **Use `var`** for zero-value initialization (e.g., `var s []int`).
- **Use `:=`** when initialization is required (e.g., `s := make([]int, 0)`).

### Map & Slice Initialization

```go
// Empty map (preferred for potential empty returns)
var m map[int]int

// Initialized map (preferred for immediate writing)
m := make(map[int]int)

// Initialized slice with capacity
s := make([]int, 0, len(input))
```

### Error Handling

In `RunTest` functions, do not panic or exit on assertion failure. Print the error and `continue` to the next test case so the developer sees a full report.

---

## 7. Helper Libraries

Use these libraries instead of standard `fmt` or `reflect` for better formatting and consistency.

### `libs/go/cmp` (Comparison)

- `cmp.EqualNumbers(a, b interface{}) bool`
- `cmp.EqualStrings(a, b string) bool`
- `cmp.EqualSlices(a, b interface{}) bool`
- `cmp.EqualMaps(a, b interface{}) bool`

### `libs/go/format` (Output)

- `format.PrintInput(map[string]interface{})`
- `format.PrintSuccess(format, args...)`
- `format.PrintFailed(format, args...)`

### `libs/go/runner` (Execution)

- `runner.InitMetrics(name)`
- `runner.ExecCountMetrics(fn, args...)`
- `runner.PrintMetrics()`

---

## 8. Dos & Don'ts

### Do

- **Do** include a complete Doc Comment with properly aligned examples.
- **Do** use `runner.ExecCountMetrics` to wrap your solution call in the test loop.
- **Do** register your solution in `main.go` immediately after creation.
- **Do** use `kebab-case` for test case keys (e.g., `"example-1-basic"`).
- **Do** implement multiple approaches if relevant, naming them `func_ApproachName`.

### Don't

- **Don't** use `fmt.Println` for test results; use `format.PrintSuccess/Failed`.
- **Don't** hardcode test inputs in the loop; use the `testCases` map.
- **Don't** leave `RunTest` functions unconnected (they must be in `main.go`).
- **Don't** use `os.Exit(1)` inside the test loop.

---

## 9. Running Solutions

From the solution directory (e.g., `problems/leetcode/go`):

```bash
# List all registered solutions
go run . -list

# Run a specific solution (by the key name registered in main.go)
go run . -solution TaskScheduler

# Short flag
go run . -s TaskScheduler
```

---

## 10. Checklist for New Solutions

Use `problems/template.go` as your base.

- [ ] **File:** Created with `snake_case` name (e.g., `621_task_scheduler.go`).
- [ ] **Doc:** Filled in all metadata fields, matching the alignment style.
- [ ] **Code:** Implemented solution in `camelCase`.
- [ ] **Test:** Implemented `RunTestPascalCase` with `continue` on failure.
- [ ] **Registry:** Added to `registerSolutions` map in `main.go` (Alphabetical order).
- [ ] **Verify:** Ran `go run . -s ProblemName` and confirmed it passes.
