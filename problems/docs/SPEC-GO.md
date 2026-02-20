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

### 3. Doc Comment

Every solution file must include a doc comment block at the top:

```go
/*
 * Problem       : <Title>
 * Topics        : <Algorithm Categories>
 * Level         : <Easy | Medium | Hard>
 * URL           : <Problem URL or "-" if none>
 * Description   : <Problem description>
 * Constraints   : <Constraints (optional)>
 * Examples      : <Examples>
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

For multiple algorithm implementations, suffix with the approach name:

```go
func findDuplicate_TortoiseHare(nums []int) int { ... }
func findDuplicate_MarkVisited(nums []int) int { ... }
func findDuplicate_HashMap(nums []int) int { ... }
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

    for name, testCase := range testCases {
        fmt.Printf("RUN %s\n", name)

        result := subarraySum(testCase.nums, testCase.k)
        format.PrintInput(map[string]interface{}{"nums": testCase.nums, "k": testCase.k})

        if !cmp.EqualNumbers(result, testCase.expect) {
            format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
            os.Exit(1)
        }
        format.PrintSuccess("test case '%s' passed", name)
    }

    fmt.Printf("\n✅ All tests passed!\n")
}
```

## Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Solution function | `camelCase` | `subarraySum`, `isBalancedBrackets` |
| Test function | `RunTest` + `PascalCase` | `RunTestSubarraySumEqualsK` |
| Constants | `PascalCase` or `UPPER_SNAKE_CASE` | `Valid4DigitPin`, `BalanceBrackets` |
| Types/Structs | `PascalCase` | `LRUCache`, `ListNode` |
| Test case names | `kebab-case` or `snake_case` | `"case-1"`, `"valid-4-digit"` |
| Multiple implementations | `functionName_ApproachName` | `findDuplicate_TortoiseHare` |

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

## Coding Style

### Variable Declaration

- Use short declaration for local variables: `result := 0`
- Group related declarations: `result, sum := 0, 0`
- Use `var` for zero-value initialization: `var passedCount uint16 = 0`

### Map Initialization

```go
// Empty map
m := map[int]int{}
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
// Empty slice
s := []int{}
s := make([]int, 0)

// With capacity hint
s := make([]rune, 0, len(input))

// With initial values
s := []int{1, 2, 3}
```

### Error Handling

Exit on test failure:

```go
if !cmp.EqualNumbers(result, testCase.expect) {
    format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
    os.Exit(1)
}
```

For non-fatal failures (continue testing):

```go
if !cmp.EqualNumbers(result, testCase.expect) {
    format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
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
- Don't use `fmt.Println` for test output (use `format` package)
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

    "github.com/ariefsibuea/algojo/libs/go/format"
    "github.com/ariefsibuea/algojo/libs/go/runner"
)

/*
 * Problem       : <Title>
 * Topics        : <Algorithm Categories>
 * Level         : <Easy | Medium | Hard>
 * URL           : <URL>
 * Description   : <Description>
 * Constraints   : <Constraints>
 * Examples      : <Examples>
 */

func RunTestXxx() {
    runner.InitMetrics("ProblemTitle")

    testCases := map[string]struct{}{
        "case-1": {},
        "case-2": {},
    }

    var passedCount uint16 = 0

    for name, testCase := range testCases {
        fmt.Printf("RUN %s\n", name)
        fmt.Println(testCase)
        // Implement test logic here

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
