def countNegatives(grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        count = 0
        row = 0
        col = cols - 1

        while row < rows and col >= 0:
            if grid[row][col] < 0:
                count += (rows - row)
                col -= 1
            else:
                row += 1
        return count
