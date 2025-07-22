from typing import Generic, List, TypeVar

T = TypeVar("T")


class Node(Generic[T]):
    def __init__(self, data: T):
        self.data = data
        self.next = None


class LinkedListStack:
    def __init__(self):
        self.top = None
        self._size = 0

    def is_empty(self) -> bool:
        """Check whether the stack is empty."""
        return self.top is None

    def __len__(self) -> int:
        """Return the number of elements in the stack."""
        return self._size

    def push(self, data: T) -> None:
        """Push an element into the stack.

        Args:
            data: The element to push.

        Time Complexity:
            O(1): Constant time to set the new node as head or top of the stack.
        """
        new_node = Node(data)
        new_node.next = self.top
        self.top = new_node
        self._size += 1

    def pop(self) -> T:
        """Remove and return the top element.

        Returns:
            T: The top element of the stack

        Raises:
            IndexError: If the stack is empty.

        Time Complexity:
            O(1): Constant time to retrieve data of the top element and move the top to the next element.
        """
        if self.is_empty():
            raise IndexError("stack is empty")

        element = self.top.data
        temp = self.top

        self.top = self.top.next
        self._size -= 1
        temp.next = None
        return element

    def peek(self) -> T:
        """Return the top element without removing from the stack.

        Returns:
            T: The top element of the stack.

        Raises:
            IndexError: If the stack is empty.

        Time Complexity:
            O(1): Constant time to retrieve data of the top element.
        """
        if self.is_empty():
            raise IndexError("stack is empty")

        return self.top.data

    def clear(self):
        """Remove all elements from the stack.

        Time Complexity:
            O(k): Where k is the current size of the stack.
        """
        node = self.top
        while node:
            temp = node
            node = node.next
            temp.next = None

        self.top = None

    def __str__(self) -> str:
        """Return a string representation of the stack."""
        if self.top is None:
            return "Stack : []"

        elements: List[str] = []
        current = self.top
        while current:
            elements.append(str(current.data))
            current = current.next

        return f"Stack: [{' -> '.join(elements)}]"


if __name__ == "__main__":
    stack = LinkedListStack()

    try:
        print("Push element 1 into stack", end="\n\n")
        stack.push(1)

        print("Push element 4 into stack", end="\n\n")
        stack.push(4)

        print("Get top element of stack")
        value = stack.peek()
        print(f"Top element of stack: {value}", end="\n\n")

        str_stack = stack.__str__()
        print(str_stack, end="\n\n")

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
