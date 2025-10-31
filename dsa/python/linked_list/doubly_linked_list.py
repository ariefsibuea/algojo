class DLLNode:
    def __init__(self, data: int):
        self.data = data
        self.next = None
        self.prev = None


class DoublyLinkedList:
    def __init__(self, head: DLLNode):
        self.head = head

    def is_empty(self) -> bool:
        return self.head is None

    def insert(self, data: int, position: int):
        """Inserts a new node to doubly linked list at a specific position.

        Time Complexity:
            O(n): In the worst case, we may need to insert a node at the end of the list.

        Space Complexity:
            O(1): Creating constant space of temporary variable.
        """
        new_node = DLLNode(data)

        if self.head is None:
            self.head = new_node
            return

        if position <= 1:
            new_node.next = self.head
            self.head.prev = new_node
            self.head = new_node
            return

        current_node = self.head
        current_position = 1

        while current_node.next and current_position < position - 1:
            current_node = current_node.next
            current_position += 1

        new_node.next = current_node.next
        current_node.next = new_node
        new_node.prev = current_node
        if new_node.next:
            new_node.next.prev = new_node

    def delete(self, position: int):
        """Delete a node at specific position in doubly linked list.

        Time Complexity:
            O(n): In the worst case, we may need to traverse the complete list of size n.

        Space Complexity:
            O(1): Creating constant space of temporary variable.
        """
        if self.head is None:
            raise ValueError("doubly linked list is empty")

        if position <= 0:
            raise ValueError("position must be greater than 0")

        if position == 1:
            temp = self.head
            self.head = self.head.next
            temp.next = None
            if self.head:
                self.head.prev = None
            return

        current_node = self.head
        current_position = 1

        while current_node.next and current_position < position - 1:
            current_node = current_node.next
            current_position += 1

        if current_node.next is None:
            raise ValueError(f"position {position} is out of range")

        temp = current_node.next
        current_node.next = current_node.next.next
        if current_node.next:
            current_node.next.prev = current_node

        temp.next = None
        temp.prev = None

    def length(self) -> int:
        """Returns length of doubly linked list by traversing all nodes.

        Time Complexity:
            O(n): Traversing the list of size n.

        Space Complexity:
            O(1): Creating constant space of temporary variable.
        """
        if self.head is None:
            raise ValueError("doubly linked list is empty")

        count = 0
        node = self.head

        while node:
            count += 1
            node = node.next

        return count

    def list(self):
        if self.head is None:
            raise ValueError("doubly linked list is empty")

        print(f"[head] = {self.head.data}")

        current = self.head.next
        current_position = 2

        while current:
            print(f"[{current_position}] = {current.data}, prev = {current.prev.data}")
            current = current.next
            current_position += 1


if __name__ == "__main__":
    print("initial doubly linked list")
    head = DLLNode(10)
    dll = DoublyLinkedList(head)
    dll.list()

    print("\ninsert 4 as head")
    dll.insert(4, 1)
    dll.list()
    print(f"length of doubly linked list = {dll.length()}")

    print("\ninsert 6 as head")
    dll.insert(6, 1)
    dll.list()
    print(f"length of doubly linked list = {dll.length()}")

    print("\ndelete node in position 2")
    dll.delete(2)
    dll.list()
    print(f"length of doubly linked list = {dll.length()}")
