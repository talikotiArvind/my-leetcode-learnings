def fairCandySwap(aliceSizes: List[int], bobSizes: List[int]) -> List[int]:
        sumA, sumB = sum(aliceSizes), sum(bobSizes)
        delta = (sumA - sumB) // 2
        bob_set = set(bobSizes)

        for a in aliceSizes:
            if (b := a - delta) in bob_set:
                return [a, b]
