def minimumSumSubarray(nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        pre = [0] * (n + 1)
        for index, value in enumerate(nums):
            pre[index + 1] = pre[index] + value
        output = float("inf")
        for size in range(l, r + 1):
            for i in range(n - size + 1):
                data = pre[i + size] - pre[i]
                if 0 < data < output:
                    output = data

        return output if output < float("inf") else -1
