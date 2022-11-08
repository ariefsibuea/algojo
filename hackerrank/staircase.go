package hackerrank

import "fmt"

func Staircase(n int32) {
	for row := n - 1; row >= 0; row-- {
		for col := int32(0); col < n; col++ {
			if col >= row {
				fmt.Printf("#")
				continue
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}
