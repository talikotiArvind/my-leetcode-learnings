def isLongPressedName(name: str, typed: str) -> bool:
        i = 0
        for index, value in enumerate(typed):
            if i < len(name) and name[i] == value:
                i += 1
            elif index == 0 or typed[index - 1] != value:
                return False
        return i == len(name)
