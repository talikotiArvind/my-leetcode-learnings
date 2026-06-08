def pivotArray(nums: List[int], pivot: int) -> List[int]:
        less = [x for x in nums if x < pivot]
        pivot_data = [x for x in nums if x == pivot]
        more = [x for x in nums if x > pivot]

        return less + pivot_data + more
