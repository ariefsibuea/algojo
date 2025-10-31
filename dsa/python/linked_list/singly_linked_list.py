class Node:
    def __init__(self, data: int):
        self.data = data
        self.next = None


class LinkedList:
    def __init__(self, head: Node):
        self.head = head

    def is_empty(self) -> bool:
        return self.head is None

    def insert(self, data: int, position: int):
        """Inserts a node at the specific position. Position 1 is the head.

        Time Complexity:
            O(n): In the worst case we may need to insert a node at the end of the list.

        Spcae Complexity:
            O(1): For creating constant space for temporary variable.
        """
        new_node = Node(data)

        if self.head is None:
            self.head = new_node
            return

        if position <= 1:
            new_node.next = self.head
            self.head = new_node
            return

        current_position = 1
        current_node = self.head

        while current_node.next and current_position < position - 1:
            current_position += 1
            current_node = current_node.next

        # NOTE: if the position > last index, then add the node after tail
        new_node.next = current_node.next
        current_node.next = new_node

    def delete(self, position: int):
        """Remove the node at the specific position.

        Time Complexity:
            O(n): In the worst case we may need to delete a node at the end of the list.

        Space Complexity:
            O(1): For constant space of temporary variable.
        """
        if self.head is None:
            raise ValueError("linked list is empty")

        if position <= 0:
            raise ValueError("position must be greater than 0")

        if position == 1:
            temp = self.head
            self.head = self.head.next
            temp.next = None
            return

        current_position = 1
        current_node = self.head

        while current_node.next and current_position < position - 1:
            current_position += 1
            current_node = current_node.next

        if current_node.next is None:
            raise ValueError(f"position {position} is out of range")

        temp = current_node.next
        current_node.next = current_node.next.next
        temp.next = None

    def length(self) -> int:
        """Returns the length of the linked list.

        Time Complexity:
            O(n): Scanning the list of size n.

        Space Complexity:
            O(1): Creating constant space of temporary variable.
        """
        if self.head is None:
            print("linked list is empty")
            return

        count = 0
        node = self.head

        while node:
            count += 1
            node = node.next

        return count

    def list(self):
        if self.head is None:
            raise ValueError("linked list is empty")

        print(f"[head] = {self.head.data}")
        current = self.head.next
        position = 2

        while current:
            print(f"[{position}] = {current.data}")
            current = current.next
            position += 1


if __name__ == "__main__":
    node = Node(1)

    print("initial linked list")
    linked_list = LinkedList(node)
    linked_list.list()

    print("\ninsert node 9 as head")
    linked_list.insert(9, 1)
    linked_list.list()
    print(f"length of linked list = {linked_list.length()}")

    print("\ninsert node 7 as head")
    linked_list.insert(7, 1)
    linked_list.list()
    print(f"length of linked list = {linked_list.length()}")

    print("\ninsert node 2")
    linked_list.insert(2, 3)
    linked_list.list()
    print(f"length of linked list = {linked_list.length()}")

    print("\ndelete head")
    linked_list.delete(1)
    linked_list.list()
    print(f"length of linked list = {linked_list.length()}")
