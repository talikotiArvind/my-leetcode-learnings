def intersection(nums1: List[int], nums2: List[int]) -> List[int]:
        output = set()
        if len(nums1) > len(nums2):
            p, q = set(nums1), set(nums2)
        else:
            p, q = set(nums2), set(nums1)
        for i in p:
            if i in q:
                output.add(i)
        return list(output)
