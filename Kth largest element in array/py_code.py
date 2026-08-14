import heapq
from typing import List


def findKthLargest(nums: List[int], k: int) -> int:
    heap = []
    for x in nums:
        heapq.heappush(heap, x)
        if len(heap) > k:
            heapq.heappop(heap)
    return heap[0]


if __name__ == "__main__":
    print(findKthLargest([3, 2, 1, 5, 6, 4], 2))        # 5
    print(findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4))  # 4