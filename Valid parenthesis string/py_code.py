def checkValidString(s):
    low, high = 0, 0
    for ch in s:
        if ch == '(':
            low += 1
            high += 1
        elif ch == ')':
            low -= 1
            high -= 1
        else:                # '*'
            low -= 1         # could be ')'
            high += 1        # could be '('
        if high < 0:         # even treating every * as '(' can't save it
            return False
        if low < 0:          # open count can't be negative — floor it
            low = 0
    return low == 0
