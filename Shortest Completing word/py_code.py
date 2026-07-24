from collections import Counter
from typing import List

class Solution:
    def shortestCompletingWord(licensePlate: str, words: List[str]) -> str:
        need = Counter(ch.lower() for ch in licensePlate if ch.isalpha())
        best = None
        for w in words:
            cnt = Counter(w)
            if all(cnt[ch] >= n for ch, n in need.items()):
                if best is None or len(w) < len(best):
                    best = w
        return best
