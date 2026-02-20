# Go Solution Specification

This document defines the structure, conventions, and coding standards for Go solutions in the algojo coding challenge repository.

## Directory Structure

```
problems/
├── leetcode/go/       # LeetCode solutions
├── hackerrank/go/     # HackerRank solutions
├── etc/go/            # Solutions from other sources
├── docs/              # Documentation
├── template.go        # Solution template
└── template.py        # Python template (reference only)
```

## File Structure

Every solution file follows this structure:

```
1. Package declaration
2. Imports
3. Doc comment (problem metadata)
4. Constants (optional)
5. Types/Structs (optional)
6. Solution function(s)
7. Test function
8. Helper functions for test input (optional)
```

### 1. Package Declaration

All solution files use `package main`:

```go
package main
```

### 2. Imports

Standard imports:

```go
import (
    "fmt"
    "os"

    "github.com/ariefsibuea/algojo/libs/go/cmp"
    "github.com/ariefsibuea/algojo/libs/go/format"
)
```

For solutions using the runner:

```go
import (
    "github.com/ariefsibuea/algojo/libs/go/format"
    "github.com/ariefsibuea/algojo/libs/go/runner"
)
```

#### Import Grouping

Imports must be organized in the following order with blank lines separating groups:

1. Standard library packages (e.g., `fmt`, `os`)
2. Project packages (e.g., `github.com/ariefsibuea/algojo/libs/go/...`)
3. Side-effect imports (e.g., `_ "path/to/package"`)

```go
import (
    "fmt"
    "os"

    "github.com/ariefsibuea/algojo/libs/go/cmp"
    "github.com/ariefsibuea/algojo/libs/go/format"
)
```

### 3. Doc Comment

Every solution file must include a doc comment block at the top:

```go
/*
 * Problem          : <Title>
 * Topics           : <Algorithm Categories>
 * Level            : <Easy | Medium | Hard>
 * URL              : <Problem URL or "-" if none>
 * Description      : <Problem description>
 * Constraints      : <Constraints (optional)>
 * Examples         : <Examples>
 */
```

#### Formatting Multi-line Content

For multi-line descriptions, constraints, or examples, align continuation lines with the previous line content using tabs:

```go
/*
 * Problem          : Task Scheduler
 * Topics           : Array, Hash Table, Greedy, Sorting, Priority Queue (Heap), Counting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/task-scheduler/
 * Description      : Schedule CPU tasks (labeled A-Z) with a cooldown constraint. Given a list of tasks and a cooldown
 * 					  period n, find the minimum time needed to execute all tasks. Each time unit can either process one
 * 					  task or be idle. The same task cannot be executed again until n time units have passed since its
 * 					  last execution. Tasks can be performed in any order, and idle periods may be necessary to satisfy
 * 					  the cooldown constraint. Return the total number of time units required.
 * Constraints      : 1 <= tasks.length <= 10^4
 * 					  tasks[i] is an uppercase English letter.
 * 					  0 <= n <= 100
 * Examples         : Example 1:
 * 					  Input: tasks = ["A","A","A","B","B","B"], n = 2
 * 					  Output: 8
 * 					  Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
 * 					  After completing task A, you must wait two intervals before doing A again.
 *
 * 					  Example 2:
 * 					  Input: tasks = ["A","C","A","B","D","B"], n = 1
 * 					  Output: 6
 * 					  Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
 */
```

#### Formatting Rules

1. **Alignment**: Continuation lines must start with `*` followed by a tab (`\t`) to align with the content start position of the previous line
2. **Blank lines**: Use ` *` (without trailing content) to separate sections within multi-line content
3. **Lists**: For constraint lists or bullet points, maintain consistent indentation with tabs

#### Tab Alignment Pattern

```
 * FieldName       : Content starts here
 * 				   continuation line aligned with tab
 * 				   another continuation line
```

The tab positions the continuation text to align with where content began on the first line (after `: `).

#### Example with All Fields

```go
/*
 * Problem          : Valid ATM PIN
 * Topics           : String, Validation
 * Level            : Easy
 * URL              : -
 * Description      : Implement a function to validate an ATM PIN with the following rules:
 * 					  - Must be 4 or 6 digits long
 * 					  - Must contain digits only
 * 					  - Must not contain repeated digits
 * 					  - Must not contain ascending or descending sequences
 * Examples         :
 * 					  Example 1:
 * 					  Input: atmpin = "1245"
 * 					  Output: 0
 *
 * 					  Example 2:
 * 					  Input: atmpin = "1234"
 * 					  Output: 5
 */
```

**Fields:**

- `Problem`: The title of the problem
- `Topics`: Algorithm/data structure categories (comma-separated)
- `Level`: Difficulty level - Easy, Medium, or Hard
- `URL`: Problem URL (use `-` if not available)
- `Description`: Brief problem description
- `Constraints`: Input constraints (optional)
- `Examples`: Input/output examples

### 4. Constants (Optional)

Define constants for return codes or fixed values:

```go
const (
    Valid4DigitPin          = 0
    Valid6DigitPin          = 1
    InvalidLength           = 2
)
```

### 5. Types/Structs (Optional)

Define data structures needed for the solution:

```go
type LRUCache struct {
    capacity int
    cache    map[int]*DLLNode
    head     *DLLNode
    tail     *DLLNode
}
```

Shared types are placed in `structures.go`:

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

### 6. Solution Function(s)

Primary solution function in `camelCase`:

```go
func subarraySum(nums []int, k int) int {
    // Implementation
}
```

For multiple algorithm implementations, append the approach name in camelCase:

```go
func findDuplicateTortoiseHare(nums []int) int { ... }
func findDuplicateMarkVisited(nums []int) int { ... }
func findDuplicateHashMap(nums []int) int { ... }
```

### 7. Test Function

Naming: `RunTest` + `PascalCase` problem name

```go
func RunTestSubarraySumEqualsK() {
    testCases := map[string]struct {
        nums   []int
        k      int
        expect int
    }{
        "case-1": {
            nums:   []int{1, 1, 1},
            k:      2,
            expect: 2,
        },
        "case-2": {
            nums:   []int{1, 2, 3},
            k:      3,
            expect: 2,
        },
    }

    for name, tc := range testCases {
        fmt.Printf("RUN %s\n", name)

        result := subarraySum(tc.nums, tc.k)
        format.PrintInput(map[string]interface{}{"nums": tc.nums, "k": tc.k})

        if !cmp.EqualNumbers(result, tc.expect) {
            format.PrintFailed("expect = %v - got = %v", tc.expect, result)
            os.Exit(1)
        }
        format.PrintSuccess("test case '%s' passed", name)
    }

    fmt.Printf("\n✅ All tests passed!\n")
}
```

#### Test Failure Messages

Test failure messages must include both the expected and actual values for easy debugging:

```go
// Good:
format.PrintFailed("expect = %v - got = %v", tc.expect, result)

// Bad: (no context)
format.PrintFailed("test failed")
```

## Naming Conventions

| Element                  | Convention                                          | Example                                    |
| ------------------------ | --------------------------------------------------- | ------------------------------------------ |
| Solution function        | `camelCase`                                         | `subarraySum`, `isBalancedBrackets`        |
| Test function            | `RunTest` + `PascalCase`                            | `RunTestSubarraySumEqualsK`                |
| Constants                | `PascalCase` (exported) or `camelCase` (unexported) | `Valid4DigitPin`, `maxRetries`             |
| Types/Structs            | `PascalCase`                                        | `LRUCache`, `ListNode`                     |
| Method receivers         | 1-2 letter abbreviation                             | `func (l *LRUCache)`, `func (n *ListNode)` |
| Test case names          | `kebab-case` or `snake_case`                        | `"case-1"`, `"valid-4-digit"`              |
| Multiple implementations | `functionName` + `ApproachName`                     | `findDuplicateTortoiseHare`                |

### Initialisms

Keep initialisms consistent in casing:

| Context    | Correct             | Incorrect           |
| ---------- | ------------------- | ------------------- |
| Exported   | `URL`, `HTTP`, `ID` | `Url`, `Http`, `Id` |
| Unexported | `url`, `http`, `id` | `uRL`, `hTTP`, `iD` |

Examples:

- `func getURL()` ✓ / `func getUrl()` ✗
- `type HTTPClient` ✓ / `type HttpClient` ✗
- `userID` ✓ / `userId` ✗

### Avoid Repetition

Avoid redundant information in names:

| Bad                  | Good           |
| -------------------- | -------------- |
| `userCountInt`       | `userCount`    |
| `userSlice`          | `users`        |
| `config.GetConfig()` | `config.Get()` |

## Test Case Structure

Use anonymous struct with map for test cases:

```go
testCases := map[string]struct {
    // Input fields
    input1 Type1
    input2 Type2
    // Expected output
    expect ExpectedType
}{
    "descriptive-name": {
        input1: value1,
        input2: value2,
        expect: expectedValue,
    },
}
```

## Helper Libraries

### cmp Package (`libs/go/cmp`)

Comparison utilities:

```go
cmp.EqualNumbers(a, b interface{}) bool    // Compare any numeric types
cmp.EqualStrings(a, b string) bool          // Compare strings
cmp.EqualBooleans(a, b bool) bool           // Compare booleans
cmp.EqualSlices(a, b interface{}) bool      // Compare slices
cmp.EqualMaps(a, b interface{}) bool        // Compare maps
cmp.IsEqual(a, b interface{}) bool          // Generic comparison
```

### format Package (`libs/go/format`)

Output formatting utilities:

```go
format.PrintInput(inputs map[string]interface{})           // Print input values
format.PrintSuccess(format string, args ...interface{})    // Print success message
format.PrintFailed(format string, args ...interface{})     // Print failure message
```

### runner Package (`libs/go/runner`)

Solution runner for CLI execution:

```go
r := runner.NewSolutionRunner()
r.RegisterSolutions(map[string]runner.TestFunc{
    "ProblemName": RunTestProblemName,
})
r.Run("ProblemName")
r.List()
```

### Package Comments for Libraries

Library packages (e.g., `libs/go/cmp`, `libs/go/format`) must include a package comment immediately above the package clause:

```go
// Package cmp provides comparison utilities for testing.
package cmp
```

## Coding Style

### Variable Names

Variable names should be proportional to scope size:

- **Small scope (1-7 lines)**: Single-letter or short names are acceptable (e.g., `i` for loop index, `n` for node)
- **Medium scope (8-15 lines)**: Single-word names (e.g., `count`, `result`, `sum`)
- **Large scope (15+ lines)**: Descriptive names (e.g., `userCount`, `processedItems`)

### Comment Style

- Complete sentences should be capitalized and end with punctuation
- Sentence fragments do not require capitalization or punctuation
- End-of-line comments for struct fields can be simple phrases

```go
// Good:
// LRUCache implements a fixed-size cache with least-recently-used eviction.
type LRUCache struct {
    capacity int  // maximum number of entries
    size     int  // current number of entries
}
```

### Variable Declaration

- Use short declaration for local variables: `result := 0`
- Group related declarations: `result, sum := 0, 0`
- Use `var` for zero-value initialization: `var passedCount uint16`

### Map Initialization

```go
// Empty map (preferred for local variables that may be returned)
var m map[int]int

// Empty map with explicit make
m := make(map[int]int)

// With initial values
m := map[int]int{0: 1}

// Type alias for readability
var bracketPair = map[rune]rune{
    '}': '{',
    ']': '[',
    ')': '(',
}
```

### Slice Initialization

```go
// Empty slice (preferred for local variables that may be returned)
var s []int

// Empty slice with explicit make
s := make([]int, 0)

// With capacity hint
s := make([]rune, 0, len(input))

// With initial values
s := []int{1, 2, 3}
```

### Error Handling

**[DEPRECATED]** Exit on test failure:

```go
if !cmp.EqualNumbers(result, tc.expect) {
    format.PrintFailed("expect = %v - got = %v", tc.expect, result)
    os.Exit(1)
}
```

For non-fatal failures (continue testing):

```go
if !cmp.EqualNumbers(result, tc.expect) {
    format.PrintFailed("expect = %v - got = %v", tc.expect, result)
    continue
}
```

## Dos & Don'ts

### Do

- Include complete doc comment with all required fields
- Use descriptive test case names
- Use helper functions for complex test input construction
- Import only necessary packages
- Use the `cmp` package for comparisons (handles type flexibility)
- Use the `format` package for consistent output
- Register solutions in the appropriate `main.go`
- Handle edge cases in test cases
- Provide multiple algorithm implementations when learning different approaches

### Don't

- Don't skip the doc comment block
- Don't use `fmt.Println`/`fmt.Printf` for test result output (use `format.PrintSuccess`, `format.PrintFailed` for consistent formatting)
- Don't hardcode test values without named constants when they represent specific codes
- Don't forget to add new solutions to `registerSolutions()` in `main.go`
- Don't use generic variable names in test structs (e.g., use `nums`, `target` instead of `a`, `b`)

## Running Solutions

From the solution directory:

```bash
# List available solutions
go run . -list

# Run specific solution
go run . -solution SubarraySumEqualsK

# Short flag
go run . -s SubarraySumEqualsK
```

## Template

Use `template.go` as the starting point for new solutions:

```go
package main

import (
    "fmt"
    "os"

    "github.com/ariefsibuea/algojo/libs/go/cmp"
    "github.com/ariefsibuea/algojo/libs/go/format"
    "github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem          : <Title>
 * Topics           : <Algorithm Categories>
 * Level            : <Easy | Medium | Hard>
 * URL              : <URL>
 * Description      : <Description>
 * 					  <continuation if needed>
 * Constraints      : <Constraints>
 * 					  - <constraint 1>
 * 					  - <constraint 2>
 * Examples         :
 * 					  Example 1:
 * 					  Input: ...
 * 					  Output: ...
 */

func solve(input type) output {
    // Implementation
}

func RunTestProblemName() {
    runner.InitMetrics("ProblemName")

    testCases := map[string]struct {
        input  inputType
        expect outputType
    }{
        "case-1": {
            input:  value1,
            expect: expected1,
        },
        "case-2": {
            input:  value2,
            expect: expected2,
        },
    }

    var passedCount uint16

    for name, tc := range testCases {
        fmt.Printf("RUN %s\n", name)

        result := solve(tc.input)
        format.PrintInput(map[string]interface{}{"input": tc.input})

        if !cmp.EqualNumbers(result, tc.expect) {
            format.PrintFailed("expect = %v - got = %v", tc.expect, result)
            os.Exit(1)
        }
        format.PrintSuccess("test case '%s' passed", name)
        passedCount++
    }

    fmt.Printf("\n📊 Test Summary: %d/%d passed\n", passedCount, len(testCases))
    runner.PrintMetrics()
}
```

## Checklist for New Solutions

- [ ] Copy `template.go` to appropriate directory with descriptive filename
- [ ] Fill in all doc comment fields
- [ ] Implement solution function(s)
- [ ] Create comprehensive test cases
- [ ] Use `cmp` package for assertions
- [ ] Use `format` package for output
- [ ] Add solution to `registerSolutions()` in `main.go`
- [ ] Run and verify all tests pass
