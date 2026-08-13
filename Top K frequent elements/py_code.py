from collections import Counter
from typing import List


def topKFrequent(nums: List[int], k: int) -> List[int]:
    counts = Counter(nums)
    buckets = [[] for _ in range(len(nums) + 1)]
    for num, freq in counts.items():
        buckets[freq].append(num)

    result = []
    for freq in range(len(buckets) - 1, 0, -1):
        for num in buckets[freq]:
            result.append(num)
            if len(result) == k:
                return result
    return result


if __name__ == "__main__":
    print(topKFrequent([1, 1, 1, 2, 2, 3], 2))  # [1, 2]
    print(topKFrequent([1], 1))                 # [1]