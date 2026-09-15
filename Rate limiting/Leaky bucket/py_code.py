import time
from collections import deque


class LeakyBucket:
    def __init__(self, capacity, leak_rate):
        self.capacity = capacity
        self.leak_rate = leak_rate
        self.queue = deque()
        self.last_leak = time.monotonic()

    def _leak(self):
        now = time.monotonic()
        elapsed = now - self.last_leak
        leaked = int(elapsed * self.leak_rate)
        for _ in range(min(leaked, len(self.queue))):
            self.queue.popleft()
        if leaked > 0:
            self.last_leak = now

    def allow(self):
        self._leak()
        if len(self.queue) < self.capacity:
            self.queue.append(time.monotonic())
            return True
        return False


if __name__ == "__main__":
    bucket = LeakyBucket(capacity=3, leak_rate=1)
    for i in range(6):
        print(i, bucket.allow())
        time.sleep(0.3)