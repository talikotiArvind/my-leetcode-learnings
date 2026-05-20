import heapq
def maximumProduct(self, nums: List[int]) -> int:
        largest = heapq.nlargest(3, nums)
        smallest = heapq.nsmallest(2, nums)
        prod1 = largest[0] * largest[1] * largest[2]
        prod2 = smallest[0] * smallest[1] * largest[0]
        return max(prod1, prod2)
