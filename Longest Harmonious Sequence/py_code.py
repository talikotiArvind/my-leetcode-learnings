# Sliding window approach

def findLHS(nums: List[int]) -> int:
        nums.sort()
        left, right = 0,0
        best = 0
        for right in range(len(nums)):
            while nums[right] - nums[left] > 1:
                left += 1
            
            if nums[right] - nums[left] == 1:
                best = max(best, right - left + 1)
        return best

# Dict (Hash Map) approach
def findLHS(nums: List[int]) -> int:
        freq = Counter(nums)
        best = 0
        for num in freq:
            if num + 1 in freq:
                best = max(best, freq[num] + freq[num + 1])
        return best
