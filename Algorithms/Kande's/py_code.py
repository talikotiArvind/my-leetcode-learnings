def max_subarray_sum(nums: list[int]) -> int:
    # Initialize both tracking variables with the first element
    max_so_far = nums[0]
    current_max = nums[0]
    
    # Iterate through the array starting from the second element
    for num in nums[1:]:
        # Decide to either extend the current subarray or start a new one
        current_max = max(num, current_max + num)
        # Update the global maximum sum found so far
        max_so_far = max(max_so_far, current_max)
        
    return max_so_far

if __name__ == "__main__":
    # Example array containing a mix of positive and negative numbers
    example_array = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    
    result = max_subarray_sum(example_array)
    print(f"Python Solution Result: {result}")  # Expected Output: 6 ([4, -1, 2, 1])
