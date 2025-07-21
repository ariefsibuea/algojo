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
        new_node = DLLNode(data)

        if self.head is None:
            self.head = new_node
            return

        if position <= 1:
            new_node.next = self.head
            self.head.prev = new_node
            self.head = new_node
