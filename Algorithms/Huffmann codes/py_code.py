import heapq
from collections import Counter

def huffman_codes(text):
    freq = Counter(text)
    if len(freq) == 1:                          # degenerate single-symbol input
        return {next(iter(freq)): "0"}
    heap = [(f, i, ch) for i, (ch, f) in enumerate(freq.items())]  # i breaks ties
    heapq.heapify(heap)
    counter = len(heap)
    while len(heap) > 1:
        lo, hi = heapq.heappop(heap), heapq.heappop(heap)
        heapq.heappush(heap, (lo[0] + hi[0], counter, (lo[2], hi[2])))
        counter += 1

    codes = {}
    def walk(node, prefix=""):
        if isinstance(node, tuple):
            walk(node[0], prefix + "0")
            walk(node[1], prefix + "1")
        else:
            codes[node] = prefix
    walk(heap[0][2])
    return codes


def decode(bits, codes):
    inv = {v: k for k, v in codes.items()}      # prefix-free => greedy works
    out, cur = [], ""
    for b in bits:
        cur += b
        if cur in inv:
            out.append(inv[cur])
            cur = ""
    return "".join(out)