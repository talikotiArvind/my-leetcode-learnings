def combinationSum(candidates: List[int], target: int) -> List[List[int]]:
        result = []

        def backtrack(start: int, current: list[int], remaining: int):
            if remaining == 0:
                result.append(current[:])
                return
            if remaining < 0:
                return

            for i in range(start, len(candidates)):
                current.append(candidates[i])
                backtrack(i, current, remaining - candidates[i])  # i not i+1, reuse allowed
                current.pop()

        backtrack(0, [], target)
        return result
