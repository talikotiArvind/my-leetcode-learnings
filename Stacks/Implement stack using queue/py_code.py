from collections import deque
class MyStack:
    def __init__(self):
        self.output = deque()

    def push(self, x: int) -> None:
        self.output.append(x)

        for _ in range(len(self.output) - 1):
            self.output.append(self.output.popleft())

    def pop(self) -> int:
        return self.output.popleft()

    def top(self) -> int:
        return self.output[0]

    def empty(self) -> bool:
        return len(self.output) == 0
