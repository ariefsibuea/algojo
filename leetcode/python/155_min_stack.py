"""
LeetCode Problem : Min Stack
Topic            : Stack, Design
Level            : Medium
URL              : https://leetcode.com/problems/min-stack/
Description      : Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
        Implement the MinStack class with methods:
        - MinStack() initializes the stack object
        - void push(int val) pushes the element val onto the stack
        - void pop() removes the element on the top of the stack
        - int top() gets the top element of the stack
        - int getMin() retrieves the minimum element in the stack
Examples         :
        Example 1:
        Input: ["MinStack","push","push","push","getMin","pop","top","getMin"]
               [[],[-2],[0],[-3],[],[],[],[]]
        Output: [null,null,null,null,-3,null,0,-2]
        Explanation: MinStack minStack = new MinStack();
                     minStack.push(-2);
                     minStack.push(0);
                     minStack.push(-3);
                     minStack.getMin(); // return -3
                     minStack.pop();
                     minStack.top();    // return 0
                     minStack.getMin(); // return -2
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
        """Initializes the stack with an empty list.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): Initializes an empty list.
        """
        self.stack = []

    def push(self, val: int) -> None:
        """Pushes an element onto the stack while tracking the minimum value.

        Args:
            val (int): The value to push onto the stack.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): Only stores one additional tuple.
        """
        if not self.stack:
            self.stack.append((val, val))
        else:
            min_value = min(val, self.stack[-1][1])
            self.stack.append((val, min_value))

    def pop(self) -> None:
        """Removes the top element from the stack.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): No additional space used.
        """
        self.stack.pop()

    def top(self) -> int:
        """Returns the top element of the stack without removing it.

        Returns:
            int: The value at the top of the stack.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): No additional space used.
        """
        return self.stack[-1][0] if self.stack else None

    def getMin(self) -> int:
        """Returns the minimum element in the stack.

        Returns:
            int: The minimum value in the stack.

        Time Complexity:
            O(1): Constant time operation.

        Space Complexity:
            O(1): No additional space used.
        """
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
