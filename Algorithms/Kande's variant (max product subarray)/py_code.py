def max_product(nums: list[int]) -> int:
    if not nums:
        return 0
        
    global_max = nums[0]
    current_max = nums[0]
    current_min = nums[0]
    
    for num in nums[1:]:
        # If the number is negative, max and min values swap potential
        if num < 0:
            current_max, current_min = current_min, current_max
            
        # Update current max and min tracker
        current_max = max(num, current_max * num)
        current_min = min(num, current_min * num)
        
        global_max = max(global_max, current_max)
        
    return global_max

if __name__ == "__main__":
    example = [2, 3, -2, 4]
    print(f"Max Product (Python): {max_product(example)}")  # Output: 6 ([2, 3])
