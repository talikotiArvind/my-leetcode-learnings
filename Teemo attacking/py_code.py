def findPoisonedDuration(timeSeries: List[int], duration: int) -> int:
        total = 0

        for i in range(len(timeSeries) - 1):
            diff = timeSeries[i + 1] - timeSeries[i]
            total += min(diff, duration)
        total += duration
        return total
