def countBits(n: int) -> List[int]:
        output = []
        for i in range(n+1):
            value = bin(i)
            output.append(value.count('1'))
        return output
