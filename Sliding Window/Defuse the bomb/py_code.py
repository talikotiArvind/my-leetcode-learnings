def decrypt(self, code: List[int], k: int) -> List[int]:
        x = len(code)
        result = [0] * x
        if k == 0:
            return result

        start = 1 if k > 0 else x + k
        end = start + abs(k)

        window_sum = sum(code[i % x] for i in range(start, end))
        result[0] = window_sum

        for i in range(1, x):
            removed = code[(start + i-1) % x]
            added = code[(end + i - 1)% x]
            window_sum = window_sum - removed + added
            result[i] = window_sum

        return result
        
