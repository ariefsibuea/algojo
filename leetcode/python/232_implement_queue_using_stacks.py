"""
LeetCode Problem : Implement Queue using Stacks
Topic            : Stack, Design, Queue
Level            : Easy
URL              : https://leetcode.com/problems/implement-queue-using-stacks/description
"""

from typing import Any


class MyQueue:

    def __init__(self):
        """Initializes the queue with two empty stacks.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): Initializes two empty lists.
        """
        self.stack_1 = []
        self.stack_2 = []

    def push(self, x: int) -> None:
        """Pushes an element to the back of the queue using stacks.

        Args:
            x (int): The element to add to the queue.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): Only adds one element to the stack.
        """
        self.stack_1.append(x)

    def pop(self) -> int:
        """Removes the element from the front of the queue and returns it.

        Returns:
            int: The element at the front of the queue.

        Time Complexity:
            O(n): Amortized O(1) per operation, but occasionally O(n) when stack_2 is empty.

        Space Complexity:
            O(1): No additional space beyond the stacks themselves.
        """
        if not self.stack_2:
            while self.stack_1:
                self.stack_2.append(self.stack_1.pop())
        return self.stack_2.pop()

    def peek(self) -> int:
        """Returns the element at the front of the queue without removing it.

        Returns:
            int: The element at the front of the queue.

        Time Complexity:
            O(n): Amortized O(1) per operation, but occasionally O(n) when stack_2 is empty.

        Space Complexity:
            O(1): No additional space beyond the stacks themselves.
        """
        if not self.stack_2:
            while self.stack_1:
                self.stack_2.append(self.stack_1.pop())
        return self.stack_2[-1]

    def empty(self) -> bool:
        """Checks if the queue is empty.

        Returns:
            bool: True if the queue is empty, False otherwise.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): No additional space used.
        """
        return max(len(self.stack_1), len(self.stack_2)) == 0


def run_tests():
    obj = MyQueue()

    res_push = obj.push(1)
    assert res_push == None, f"expected None, got {res_push}"

    res_push = obj.push(2)
    assert res_push == None, f"expected None, got {res_push}"

    res_peek = obj.peek()
    assert res_peek == 1, f"expected 1, got {res_peek}"

    res_pop = obj.pop()
    assert res_pop == 1, f"expected 1, got {res_pop}"

    res_empty = obj.empty()
    assert res_empty == False, f"expected False, got {res_empty}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
