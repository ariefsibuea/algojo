package main

import "fmt"

/**
 * Towers of Hanoi
 *
 * The Towers of Hanoi is a mathematical puzzle. It consist of three pegs or towers and a number of disks of different
 * sizes which can slude onto any tower. The puzzle starts with the disks on one tower in ascending order of size, the
 * smallest at the top, thus making a conical shape. The objective of the puzzle is to move the entire stack to another
 * tower, satisfying the following rules:
 * 		- Only one disk may be moved at a time.
 * 		- Each move consist of taking the upper disk from one of the tower and sliding it onto another tower, on top of
 * 			the other disks that may already be present on that tower.
 * 		- No disk may be placed on top of a smaller disk.
 *
 * Algorithm:
 * 		- Move the top n-1 disks from Source to Auxiliary tower.
 * 		- Move the nth disk from Source to Destination tower.
 * 		- Move the n-1 disks from Auxiliary tower to Destination tower.
 * Transferring the top n-1 disks from Source to Auxiliary tower can again be thought of as a fresh problem and can be
 * solved in the same manner. Once we solve Towers of Hanoi with three disks, we can solve it with any number of disks
 * with the above algorithm.
 */

func TowersOfHanoi(n int, from, to, aux string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from tower %s to tower %s\n", from, to)
		return
	}

	// move top n-1 disks from 'from' to 'aux' using 'to' as auxiliary
	TowersOfHanoi(n-1, from, aux, to)

	// move remaining disks from 'from' to 'to'
	fmt.Printf("Move disk %d from tower %s to tower %s\n", n, from, to)

	// move n-1 disks from 'aux' to 'to' using 'from' as auxiliary
	TowersOfHanoi(n-1, aux, to, from)
}
