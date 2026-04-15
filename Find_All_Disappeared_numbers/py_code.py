def findDisappearedNumbers(nums: List[int]) -> List[int]:
        output = []
        numbers = set(nums)
        for i in range(1, len(nums)+1):
            if i not in numbers:
                output.append(i)
        return output
