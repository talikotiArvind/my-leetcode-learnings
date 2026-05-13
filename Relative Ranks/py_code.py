def findRelativeRanks(score: List[int]) -> List[str]:
        ranked = sorted(score, reverse=True)
        rank = {}
        
        for i, s in enumerate(ranked):
            match i:
                case 0:
                    rank[s] = "Gold Medal"
                case 1:
                    rank[s] = "Silver Medal"
                case 2:
                    rank[s] = "Bronze Medal"
                case _:
                    rank[s] = str(i + 1)
        return [rank[s] for s in score]
