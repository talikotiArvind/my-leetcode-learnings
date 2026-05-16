from collections import Counter
def intersect(nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        count = Counter(nums1)
        result = []

        for n in nums2:
            if count[n] > 0:
                result.append(n)
                count[n] -= 1

        return result
