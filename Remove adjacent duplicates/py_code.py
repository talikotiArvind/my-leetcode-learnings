def removeDuplicates(s):
    stack = []
    for ch in s:
        if stack and stack[-1] == ch:
            stack.pop()          # cancel the matching pair
        else:
            stack.append(ch)     # push
    return "".join(stack)
