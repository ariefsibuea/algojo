class DLLNode:
    def __init__(self, key: int = 0, value: int = 0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.cache = {}  # map key to doubly linked list node
        self.size = 0
        self.capacity = capacity

        self.head = DLLNode()  # pointing to head -> the most recently used
        self.tail = DLLNode()  # pointing to tail -> the least recently used
        self.head.next = self.tail
        self.tail.prev = self.head

    def _move_to_head(self, node: DLLNode):
        self._remove(node)
        self._insert(node)

    def _remove(self, node: DLLNode):
        prev = node.prev
        new = node.next
        prev.next = new
        new.prev = prev

    def _insert(self, node: DLLNode):
        # always add node to the right after head
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node

    def _pop(self):
        # always pop the tail (the least recently used)
        result = self.tail.prev
        self._remove(result)
        return result

    def get(self, key: int) -> int:
        node = self.cache.get(key, None)
        if not node:
            return -1

        # move the hit node to head
        self._move_to_head(node)
        return node.value

    def put(self, key: int, value: int) -> None:
        node = self.cache.get(key, None)

        if not node:
            new_node = DLLNode(key, value)
            self.cache[key] = new_node
            self._insert(new_node)
            self.size += 1

            if self.size > self.capacity:
                tail = self._pop()
                del self.cache[tail.key]
                self.size -= 1
        else:
            node.value = value
            self._move_to_head(node)


if __name__ == "__main__":
    """
    Input:
        ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
        [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

    Output
        [null, null, null, 1, null, -1, null, -1, 3, 4]
    """

    cache = LRUCache(2)

    cache.put(1, 1)
    print(f"action = put(1, 1)")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    cache.put(2, 2)
    print(f"action = put(2, 2)")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    result = cache.get(1)
    print(f"action = get(1)")
    print(f"result = {result}")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    cache.put(3, 3)
    print(f"action = put(3, 3)")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    result = cache.get(2)
    print(f"action = get(2)")
    print(f"result = {result}")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    cache.put(4, 4)
    print(f"action = put(4, 4)")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    result = cache.get(1)
    print(f"action = get(1)")
    print(f"result = {result}")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    result = cache.get(3)
    print(f"action = get(3)")
    print(f"result = {result}")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")

    result = cache.get(4)
    print(f"action = get(4)")
    print(f"result = {result}")
    print(f"most recently used = {cache.head.next.value}")
    print(f"least recently used = {cache.tail.prev.value}", end="\n\n")
