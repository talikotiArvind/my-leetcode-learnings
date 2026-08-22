import random

def reservoir_sample(stream, k):
    res = []
    for i, item in enumerate(stream):
        if i < k:
            res.append(item)
        else:
            j = random.randint(0, i)      # inclusive on both ends
            if j < k:
                res[j] = item
    return res