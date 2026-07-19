def merge(intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) <= 1:
            return intervals

        # sort by start time
        intervals.sort(key=lambda x: x[0])

        res = [intervals[0]]

        for cur in intervals[1:]:
            last = res[-1]
            if cur[0] <= last[1]:
                # overlap — extend the end of the last merged interval
                last[1] = max(last[1], cur[1])
            else:
                # no overlap — start a new interval
                res.append(cur)

        return res
