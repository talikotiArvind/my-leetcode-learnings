def targetIndices(nums: List[int], target: int) -> List[int]:
        return [i for i, value in enumerate(sorted(nums)) if value == target]
