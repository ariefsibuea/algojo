def tower_of_hanoi(n: int, from_peg: str, to_peg: str, aux_peg: str):
    """Solves the Tower of Hanoi problem using recursion.

    Problem:
        The Tower of Hanoi is a classic puzzle where you have three pegs and n disks of different sizes which can slide
        onto any peg. The puzzle starts with the disks neatly stacked in ascending order of size on one peg, the
        smallest at the top, making a conical shape. The objective is to move the entire stack to another peg, obeying
        the following rules:
            1. Only one disk may be moved at a time.
            2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack
            or on an empty peg.
            3. No disk may be placed on top of a smaller disk.

    Solution Approach:
        This function uses a recursive approach to solve the problem.
            - Move the top n-1 disks from the source peg to the auxiliary peg.
            - Move the nth (largest) disk from the source peg to the destination peg.
            - Move the n-1 disks from the auxiliary peg to the destination peg.

    Args:
        n (int): Number of disks to move.
        from_peg (str): The name or label of the source peg.
        to_peg (str): The name or label of the destination peg.
        aux_peg (str): The name or label of the auxiliary peg used for temporary storage.

    Returns:
        None. The function prints each move required to solve the puzzle.

    Time Complexity:
        O(2^n): Where n is the number of disks. Each disk move results in two recursive calls for n-1 disks, leading to
        exponential growth.

    Space Complexity:
        O(n): Due to the recursion stack depth.
    """
    if n == 1:
        print(f"Move disk 1 from peg {from_peg} to peg {to_peg}")
        return

    # Move top n-1 disks from A to B using C as auxiliary
    tower_of_hanoi(n-1, from_peg, aux_peg, to_peg)

    print(f"Move disk {n} from peg {from_peg} to peg {to_peg}")

    # Move n-1 disks from B to C using A as auxiliary
    tower_of_hanoi(n-1, aux_peg, to_peg, from_peg)

if __name__ == "__main__":
    tower_of_hanoi(4, "A", "C", "B")
