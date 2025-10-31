package main

import "fmt"

// Binary generates all the binaries of 'n' bits with 'arrBinary' is an array of size n.
func Binary(n int, arrBinary []int) {
	if n < 1 {
		fmt.Printf("%v\n", arrBinary)
		return
	}

	arrBinary[n-1] = 0
	Binary(n-1, arrBinary)
	arrBinary[n-1] = 1
	Binary(n-1, arrBinary)
}
