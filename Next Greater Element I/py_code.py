def nextGreaterElement(nums1: list[int], nums2: list[int]) -> list[int]:
    nge_map = {}
    stack = []

    for num in nums2:
        while stack and num > stack[-1]:
            top = stack.pop()
            nge_map[top] = num
        stack.append(num)

    while stack:
        top = stack.pop()
        nge_map[top] = -1

    return [nge_map[num] for num in nums1]
