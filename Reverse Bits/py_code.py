def reverseBits(n: int) -> int:
        return int(f"{n:032b}"[::-1], 2)
