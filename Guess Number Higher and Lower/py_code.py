def guessNumber(n: int) -> int:
        left, right = 0, n
        while left <= right:
            middle = (left + right) // 2
            match guess(middle):
                case 0:
                    return middle
                case -1:
                    right = middle - 1
                case 1:
                    left = middle + 1
