// This is the standard template for input/output.
package template

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func temp_main() {
	defer writer.Flush()

	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("%d\n", a+b)
}
