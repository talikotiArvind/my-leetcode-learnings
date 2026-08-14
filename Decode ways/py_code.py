def numDecodings(s: str) -> int:
    if not s or s[0] == "0":
        return 0

    prev, curr = 1, 1  # ways up to i-2, i-1
    for i in range(1, len(s)):
        temp = 0
        if s[i] != "0":
            temp += curr
        if 10 <= int(s[i - 1:i + 1]) <= 26:
            temp += prev
        prev, curr = curr, temp
    return curr


if __name__ == "__main__":
    print(numDecodings("12"))    # 2  -> "AB", "L"
    print(numDecodings("226"))   # 3  -> "BZ", "VF", "BBF"
    print(numDecodings("06"))    # 0