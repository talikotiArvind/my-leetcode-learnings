def findTheDistanceValue(arr1: List[int], arr2: List[int], d: int) -> int:
  arr1.sort(),arr2.sort()
  i = j = 0
  count = 0 
  while i < len(arr1) and j < len(arr2):
      if abs(arr1[i] - arr2[j]) <= d:
          i += 1
      elif arr1[i] > arr2[j]:
          j += 1
      else:
          i += 1
          count += 1
  if i < len(arr1):
      count += len(arr1) - i
  return count 
