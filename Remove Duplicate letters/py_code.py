def removeDuplicateLetters(s):
    last = {c: i for i, c in enumerate(s)}   # last index each char appears
    stack = []
    seen = set()
    for i, c in enumerate(s):
        if c in seen:
            continue
        # pop bigger chars that still appear later
        while stack and stack[-1] > c and last[stack[-1]] > i:
            seen.discard(stack.pop())
        stack.append(c)
        seen.add(c)
    return "".join(stack)
