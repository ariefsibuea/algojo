from typing import Generic, TypeVar


T = TypeVar("T")


class ArrayQueue(Generic[T]):
    def __init__(self, capacity: int = 10):
        self.front = -1
        self.rear = -1
        self.capacity = capacity
        self.array = [None] * capacity

    def is_empty(self) -> bool:
        return self.front == -1

    def is_full(self) -> bool:
        return (self.rear + 1) % self.capacity == self.front

    def size(self) -> int:
        if self.front == -1:
            return 0
        elif self.front <= self.rear:
            return self.rear - self.front + 1
        else:
            return (self.capacity - self.front + self.rear + 1) % self.capacity

    def enqueue(self, data: T):
        if self.is_full():
            raise OverflowError("queue overflow")

        self.rear = (self.rear + 1) % self.capacity
        self.array[self.rear] = data
        if self.front == -1:
            self.front = self.rear

    def dequeue(self):
        if self.is_empty():
            raise IndexError("queue is empty")

        data = self.array[self.front]
        if self.front == self.rear:
            self.front = self.rear = -1
        else:
            self.front = (self.front + 1) % self.capacity

        return data
