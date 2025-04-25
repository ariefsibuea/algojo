"""
LeetCode Problem : Min Stack
Topic            : Stack, Design
Level            : Medium
URL              : https://leetcode.com/problems/min-stack/description/
"""


class MinStack:

    """
    NOTE:
        Below is my previous solution. There are some issues:
            1. Overcomplicated logic to manage size and top_index.
            2. Inefficient pop, setting self.stack[self.top_index] = None does not reduce memory usage.
            3. Fragile state management, mixing size and top_index increases the risk of bugs e.g., invalid indices.
    """
    # def __init__(self):
    #     self.stack = []
    #     self.size = 0
    #     self.top_index = -1

    # def push(self, val: int) -> None:
    #     if not self.stack:
    #         self.stack.append({"value": val, "minimum": val})
    #         self.size += 1
    #         self.top_index += 1
    #     elif (self.size - 1) > self.top_index:
    #         min_value = val
    #         if self.top_index > -1 and self.stack[self.top_index]["minimum"] < val:
    #             min_value = self.stack[self.top_index]["minimum"]

    #         self.stack[self.top_index + 1] = {"value": val, "minimum": min_value}
    #         self.top_index += 1
    #     else:
    #         min_value = val
    #         if self.stack[self.top_index]["minimum"] < val:
    #             min_value = self.stack[self.top_index]["minimum"]

    #         self.stack.append({"value": val, "minimum": min_value})
    #         self.size += 1
    #         self.top_index += 1

    # def pop(self) -> None:
    #     self.stack[self.top_index] = None
    #     if self.top_index > -1:
    #         self.top_index -= 1

    # def top(self) -> int:
    #     if self.top_index > -1:
    #         return self.stack[self.top_index]["value"]
    #     return 0

    # def getMin(self) -> int:
    #     if self.top_index > -1:
    #         return self.stack[self.top_index]["minimum"]
    #     return 0

    def __init__(self):
        self.stack = []

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append((val, val))
        else:
            min_value = min(val, self.stack[-1][1])
            self.stack.append((val, min_value))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0] if self.stack else None

    def getMin(self) -> int:
        return self.stack[-1][1] if self.stack else None


def run_tests():
    min_stack = MinStack()

    min_stack.push(6)
    min_stack.push(6)
    min_stack.push(7)

    top = min_stack.top()
    assert top == 7, f"top: expected 7, got {top}"

    min_stack.pop()

    min = min_stack.getMin()
    assert min == 6, f"getMin: expected 6, got {min}"

    min_stack.pop()

    min = min_stack.getMin()
    assert min == 6, f"getMin: expected 6, got {min}"

    min_stack.pop()
    min_stack.push(7)

    top = min_stack.top()
    assert top == 7, f"top: expected 7, got {top}"

    min = min_stack.getMin()
    assert min == 7, f"getMin: expected 7, got {min}"

    min_stack.push(-8)

    top = min_stack.top()
    assert top == -8, f"top: expected -8, got {top}"

    min = min_stack.getMin()
    assert min == -8, f"getMin: expected -8, got {min}"

    min_stack.pop()

    min = min_stack.getMin()
    assert min == 7, f"getMin: expected 7, got {min}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
