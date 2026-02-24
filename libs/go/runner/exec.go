package runner

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type MetricsCollector struct {
	executions  []ExecutionMetrics
	name        string
	totalTime   time.Duration
	totalMemory uint64
	totalAllocs uint64
}

type ExecutionMetrics struct {
	ExecutionTime time.Duration
	MemoryUsed    uint64
	AllocsCount   uint64
}

var globalCollector *MetricsCollector

func InitMetrics(name string) {
	globalCollector = &MetricsCollector{
		executions: make([]ExecutionMetrics, 0),
		name:       name,
	}
}

func ExecCountMetrics(fn interface{}, args ...interface{}) interface{} {
	var metrics ExecutionMetrics
	var m1, m2 runtime.MemStats

	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		panic("first argument must be a function")
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	runtime.GC()
	runtime.ReadMemStats(&m1)

	startTime := time.Now()
	results := fnValue.Call(in)
	metrics.ExecutionTime = time.Since(startTime)

	runtime.ReadMemStats(&m2)
	metrics.MemoryUsed = m2.TotalAlloc - m1.TotalAlloc
	metrics.AllocsCount = m2.Mallocs - m1.Mallocs

	if globalCollector != nil {
		globalCollector.executions = append(globalCollector.executions, metrics)
		globalCollector.totalTime += metrics.ExecutionTime
		globalCollector.totalMemory += metrics.MemoryUsed
		globalCollector.totalAllocs += metrics.AllocsCount
	}

	switch len(results) {
	case 0:
		return nil
	case 1:
		return results[0].Interface()
	default:
		out := make([]interface{}, len(results))
		for i, r := range results {
			out[i] = r.Interface()
		}
		return out
	}
}

func PrintMetrics() {
	if globalCollector == nil || len(globalCollector.executions) == 0 {
		fmt.Println("No metrics collected")
		return
	}

	// Helper function to print a row with proper alignment
	printRow := func(label, value string, indent int) {
		// Build the content with indent
		indentStr := ""
		if indent > 0 {
			indentStr = fmt.Sprintf("%*s", indent, "")
		}

		// Use fixed-width formatting: total width is 58 chars
		content := fmt.Sprintf("%s%-*s", indentStr, 58-indent, label+" "+value)
		fmt.Printf("║ %s ║\n", content)
	}

	execCount := len(globalCollector.executions)

	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ %-58s ║\n", "Performance Metrics: "+globalCollector.name)
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-58s ║\n", fmt.Sprintf("Total Executions: %d", execCount))
	fmt.Println("╠════════════════════════════════════════════════════════════╣")

	// Execution Time
	avgTime := globalCollector.totalTime / time.Duration(execCount)
	minTime, maxTime := getMinMaxTime(globalCollector.executions)

	fmt.Printf("║ %-58s ║\n", "Execution Time:")
	printRow("Total:", formatDuration(globalCollector.totalTime), 2)
	printRow("Average:", formatDuration(avgTime), 2)
	printRow("Min:", formatDuration(minTime), 2)
	printRow("Max:", formatDuration(maxTime), 2)

	// Memory Usage
	avgMemory := globalCollector.totalMemory / uint64(execCount)
	minMemory, maxMemory := getMinMaxMemory(globalCollector.executions)

	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-58s ║\n", "Memory Usage:")
	printRow("Total:", formatBytes(globalCollector.totalMemory), 2)
	printRow("Average:", formatBytes(avgMemory), 2)
	printRow("Min:", formatBytes(minMemory), 2)
	printRow("Max:", formatBytes(maxMemory), 2)

	// Allocations
	avgAllocs := globalCollector.totalAllocs / uint64(execCount)
	minAllocs, maxAllocs := getMinMaxAllocs(globalCollector.executions)

	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-58s ║\n", "Allocations:")
	printRow("Total:", fmt.Sprintf("%d", globalCollector.totalAllocs), 2)
	printRow("Average:", fmt.Sprintf("%d", avgAllocs), 2)
	printRow("Min:", fmt.Sprintf("%d", minAllocs), 2)
	printRow("Max:", fmt.Sprintf("%d", maxAllocs), 2)
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

func ResetMetrics() {
	if globalCollector != nil {
		globalCollector.executions = make([]ExecutionMetrics, 0)
		globalCollector.totalTime = 0
		globalCollector.totalMemory = 0
		globalCollector.totalAllocs = 0
	}
}

func getMinMaxTime(executions []ExecutionMetrics) (time.Duration, time.Duration) {
	if len(executions) == 0 {
		return 0, 0
	}

	minTime, maxTime := executions[0].ExecutionTime, executions[0].ExecutionTime
	for _, e := range executions[1:] {
		minTime = min(minTime, e.ExecutionTime)
		maxTime = max(maxTime, e.ExecutionTime)
	}

	return minTime, maxTime
}

func getMinMaxMemory(executions []ExecutionMetrics) (uint64, uint64) {
	if len(executions) == 0 {
		return 0, 0
	}

	minMem, maxMem := executions[0].MemoryUsed, executions[0].MemoryUsed
	for _, e := range executions[1:] {
		minMem = min(minMem, e.MemoryUsed)
		maxMem = max(maxMem, e.MemoryUsed)
	}

	return minMem, maxMem
}

func getMinMaxAllocs(executions []ExecutionMetrics) (uint64, uint64) {
	if len(executions) == 0 {
		return 0, 0
	}

	minAllocs, maxAllocs := executions[0].AllocsCount, executions[0].AllocsCount
	for _, e := range executions[1:] {
		minAllocs = min(minAllocs, e.AllocsCount)
		maxAllocs = max(maxAllocs, e.AllocsCount)
	}

	return minAllocs, maxAllocs
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2f µs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2f ms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2f s", d.Seconds())
}

func formatBytes(b uint64) string {
	const unit = 1024

	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
