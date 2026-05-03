def searchRange(nums: List[int], target: int) -> List[int]:
        def binsearch(find_first: bool):
            l, r, result = 0, len(nums) - 1, -1
            while l <= r:
                m = (l+r)//2
                if nums[m] == target:
                    result = m
                    if find_first:
                        r = m - 1
                    else:
                        l = m + 1
                elif nums[m] < target:
                    l = m + 1
                else:
                    r = m - 1
            return result
        return [binsearch(True), binsearch(False)]
