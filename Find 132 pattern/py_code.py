def find132pattern(nums: List[int]) -> bool:
        stack = []
        data = float('-inf')
        for n in reversed(nums):
            if n < data:
                return True
            while stack and stack[-1] < n:
                data = stack.pop()
            stack.append(n)
        return False
