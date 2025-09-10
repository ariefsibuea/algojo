package main

import "fmt"

func main() {
	// TowersOfHanoi(3, "A", "B", "C")

	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Is array %v in ascending order? %v\n", nums, IsArrayInAscOrder(5, nums))
	fmt.Printf("Is array %v in descending order? %v\n", nums, IsArrayInDescOrder(5, nums))
}
