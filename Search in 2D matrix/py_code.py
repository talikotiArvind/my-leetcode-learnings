def searchMatrix(matrix: List[List[int]], target: int) -> bool:
      x, y = len(matrix), len(matrix[0])
      low, high = 0, x * y -1
      while low <= high:
          mid = (low + high) // 2
          value = matrix[mid // y][mid % y]
          if value == target:
              return True
          elif value < target:
              low = mid + 1
          else:
              high = mid - 1
      return False
