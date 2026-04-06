def moveZeroes(nums) -> None:
    """
    Do not return anything, modify nums in-place instead.
    """
    k = 0
    for i in range(len(nums)):
        if nums[i] != 0:
            nums[k], nums[i] = nums[i], nums[k]
            k+=1

nums = [0,1,0,3,12]
moveZeroes(nums)
print(nums)