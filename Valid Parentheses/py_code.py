def isValid(s: str) -> bool:
    stack = []
    pairs = {')': '(', ']': '[', '}': '{'}
    for ch in s:
        if ch in pairs:
            if not stack or stack.pop() != pairs[ch]:
                return False
        else:
            stack.append(ch)
    return not stack


if __name__ == "__main__":
    print(isValid("()[]{}"))  # True
    print(isValid("(]"))      # False
    print(isValid("([)]"))    # False
    print(isValid("{[]}"))    # True
