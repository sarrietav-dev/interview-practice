class MyQueue:
    def __init__(self):
        self.stack = []
        self.temp_stack = []

    def push(self, x: int) -> None:
        for _ in range(len(self.stack)):
            self.temp_stack.append(self.stack.pop())
        self.stack.append(x)
        for _ in range(len(self.temp_stack)):
            self.stack.append(self.temp_stack.pop())

    def pop(self) -> int:
        return self.stack.pop()

    def peek(self) -> int:
        return self.stack[-1]

    def empty(self) -> bool:
        return len(self.stack) == 0
