import time
from collections import deque


class SlidingWindowCounter:
    def __init__(self, limit, window_seconds):
        self.limit = limit
        self.window = window_seconds
        self.timestamps = deque()

    def allow(self):
        now = time.monotonic()
        while self.timestamps and now - self.timestamps[0] > self.window:
            self.timestamps.popleft()
        if len(self.timestamps) < self.limit:
            self.timestamps.append(now)
            return True
        return False


if __name__ == "__main__":
    limiter = SlidingWindowCounter(limit=3, window_seconds=1)
    for i in range(6):
        print(i, limiter.allow())
        time.sleep(0.3)