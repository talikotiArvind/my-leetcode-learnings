def flipAndInvertImage(image: List[List[int]]) -> List[List[int]]:
        for row in image:
            low, high = 0, len(row) - 1
            while low < high:
                row[low], row[high] = row[high] ^ 1, row[low] ^ 1
                low += 1
                high -= 1
            if low == high:
                row[low] ^= 1
        return image
