class Node:
    def __init__(self, data: int):
        self.data = data
        self.next = None


class LinkedList:
    def __init__(self, head: Node):
        self.head = head

    def is_empty(self):
        return self.head is None

    def insert(self, node: Node, position: int):
        """Inserts a node at the specific position. Position 1 is the head.

        Time Complexity:
            O(n): In the worst case we may need to insert a node at the end of the list.

        Spcae Complexity:
            O(1): For creating one temporary variable.
        """
        if self.head is None:
            self.head = node
            return

        if position == 1:
            node.next = self.head
            self.head = node
            return

        current_position = 1
        current_node = self.head

        while current_node.next and current_position < position - 1:
            current_position += 1
            current_node = node.next

        # NOTE: if the position > last index, then add the node after tail
        node.next = current_node.next
        current_node.next = node

    def delete(self, position: int):
        """Remove the node at the specific position.

        Time Complexity:
            O(n): In the worst case we may need to delete a node at the end of the list.

        Space Complexity:
            O(1): For one temporary variable.
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

        if position - 1 > current_position:
            raise ValueError(f"position {position} is out of range")

        if current_node.next is None:
            print(f"position {position} has been empty")

        temp = current_node.next
        current_node.next = current_node.next.next
        temp.next = None

    def length(self) -> int:
        """Returns the length of the linked list.

        Time Complexity:
            O(n): Scanning the list of size n.

        Space Complexity:
            O(1): Creating a temporary variable.
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

    def print_nodes(self):
        if self.head is None:
            raise ValueError("linked list is empty")

        print(f"head = {self.head.data}")
        current = self.head
        position = 1

        while current.next:
            current = current.next
            position += 1
            print(f"position {position} = {current.data}")


if __name__ == "__main__":
    node = Node(1)

    print("initial linked list")
    linked_list = LinkedList(node)
    linked_list.print_nodes()

    print("\ninsert node 2")
    linked_list.insert(Node(2), 3)
    linked_list.print_nodes()
    print(f"length of linked list = {linked_list.length()}")

    print("\ndelete head")
    linked_list.delete(1)
    linked_list.print_nodes()
    print(f"length of linked list = {linked_list.length()}")
