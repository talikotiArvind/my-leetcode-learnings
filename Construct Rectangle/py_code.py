def constructRectangle(area: int) -> List[int]:
        value = int(area ** (0.5))
        while area % value != 0:
            value -= 1
        return [area // value, value]
