def containsNearbyDuplicate(nums: List[int], k: int) -> bool:
      if len(nums) == len(set(nums)):
          return False
      seen = {}
      for index, value in enumerate(nums):
          if value in seen and index - seen[value] <= k:
              return True
          seen[value] = index
      return False
