def maxArea(height: List[int]) -> int:
        left, right = 0, len(height) - 1
        max_value = 0
        while left < right:
            width = right - left
            water = width * min(height[left], height[right])
            max_value = max(max_value, water)

            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return max_value
        
