def get_sum(a: int, b: int) -> int:
    mask = 0xFFFFFFFF
    while b & mask:
        carry = (a & b) << 1
        a = (a ^ b) & mask
        b = carry & mask
    result = a & mask
    if result > 0x7FFFFFFF:
        result = ~(result ^ mask)
    return result


if __name__ == "__main__":
    assert get_sum(1, 2) == 3
    assert get_sum(-2, 3) == 1
    assert get_sum(-12, 8) == -4
    assert get_sum(0, 0) == 0
    assert get_sum(-5, -7) == -12
    print("All tests passed.")
