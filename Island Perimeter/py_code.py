def islandPerimeter(grid: List[List[int]]) -> int:
        row, column = len(grid), len(grid[0])
        result = 0
        for x in range(row):
            for y in range(column):
                if grid[x][y] == 1:
                    result += 4
                    if x > 0 and grid[x-1][y] == 1:
                        result -= 2
                    if y > 0 and grid[x][y-1] == 1:
                        result -= 2
        return result
