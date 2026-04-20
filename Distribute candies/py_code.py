def distributeCandies(candyType: List[int]) -> int:
        return min(len(candyType) // 2, len(set(candyType)))
