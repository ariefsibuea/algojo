from typing import Generic, TypeVar

T = TypeVar("T")


class ArrayStack(Generic[T]):
    def __init__(self, capacity: int = 10):
        self.top = -1
        self.capacity = capacity
        self.array = [None] * capacity

    def is_empty(self) -> bool:
        """Check whether the stack is empty."""
        return self.top == -1

    def is_full(self) -> bool:
        """Check whether the stack is full."""
        return self.top == self.capacity - 1

    def push(self, data: T) -> None:
        """Push an element into the stack.

        Args:
            data: The element to push.

        Raises:
            OverflowError: If the stack is full.

        Time Complexity:
            O(1): Directly set the element to specific slot by tracked index.
        """
        if self.is_full():
            raise OverflowError("stack overflow")

        self.top += 1
        self.array[self.top] = data

    def pop(self) -> T:
        """Remove and return the top element.

        Returns:
            T: The top element of the stack

        Raises:
            IndexError: If the stack is empty.

        Time Complexity:
            O(1): Directly return the top element retrived using tracked index.
        """
        if self.is_empty():
            raise IndexError("stack is empty")

        element = self.array[self.top]
        self.array[self.top] = None
        self.top -= 1
        return element

    def peek(self) -> T:
        """Return the top element without removing from the stack.

        Returns:
            T: The top element of the stack.

        Raises:
            IndexError: If the stack is empty.

        Time Complexity:
            O(1): Directly return the top element retrived using tracked index.
        """
        if self.is_empty():
            raise IndexError("stack is empty")

        return self.array[self.top]

    def clear(self):
        """Remove all elements from the stack.

        Time Complexity:
            O(k): Where k is the current size of the stack.
        """
        self.array = [None] * (self.top + 1)
        self.top = -1

    def size(self):
        """Return the number of elements in the stack."""
        return self.top + 1


if __name__ == "__main__":
    stack = ArrayStack[int]()

    try:
        print("Push element 1 into stack", end="\n\n")
        stack.push(1)

        print("Push element 4 into stack", end="\n\n")
        stack.push(4)

        print("Get top element of stack")
        value = stack.peek()
        print(f"Top element of stack: {value}", end="\n\n")

        print("Pop element from stack")
        value = stack.pop()
        print(f"Popped value: {value}", end="\n\n")

        print("Pop element from stack")
        value = stack.pop()
        print(f"Popped value: {value}", end="\n\n")

        print("Pop element from stack")  # stack is empty
        value = stack.pop()
        print(f"Popped value: {value}")
    except Exception as e:
        print(f"Exception: {e}")
