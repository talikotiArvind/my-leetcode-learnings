def canJump(nums: List[int]) -> bool:
        farthest = 0
        last = len(nums) - 1
        for index, element in enumerate(nums):
            if index > farthest:
                return False
            farthest = max(farthest, index + element)

            if farthest >= last:
                return True
        return fartheest >= last
  
