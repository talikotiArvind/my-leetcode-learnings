def findShortestSubArray(nums: List[int]) -> int:
        first = {}
        count = defaultdict(int)
        degree = 0
        result = 0

        for i, x in enumerate(nums):
            if x not in first:
                first[x] = i
            count[x] += 1

            if count[x] > degree:
                degree = count[x]
                result = i - first[x] + 1
            elif count[x] == degree:
                result = min(result, i - first[x] + 1)

        return result
